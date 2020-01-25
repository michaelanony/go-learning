package session

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"sync"
)

type RedisSession struct {
	sessionId string
	pool *redis.Pool
	//设置session，可以先放在内存的map中
	//批量导入redis，提升性能
	sessionMap map[string]interface{}
	//读写锁
	rwLock sync.RWMutex
	flag int
}
//用常量定义状态
const(
	//内存数据没变化
	SessionFlagNone = iota
	//有变化
	SessionFlagModify
)
func NewRedisSession(id string,pool *redis.Pool) *RedisSession{
	s := &RedisSession{
		sessionId:  id,
		pool:       pool,
		sessionMap: make(map[string]interface{},16),
		flag:       SessionFlagNone,
	}
	return s
}

func (r *RedisSession)Set(key string,value interface{}) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	//Set value
	r.sessionMap[key] = value
	r.flag = SessionFlagModify
	return
}
func (r *RedisSession)Get(key string) (ret interface{},err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	//先看内存有没有数据
	ret,ok := r.sessionMap[key]
	if !ok{
		err = errors.New("key not exists")
	}
	return
}
func (r *RedisSession)LoadFromRedis (ret interface{},err error)  {
	conn := r.pool.Get()
	reply,err :=conn.Do("GET",r.sessionId)
	if err!=nil{
		return
	}
	data,err := redis.String(reply,err)
	if err!=nil{
		return
	}
	//取到的东西反序列化内存的map
	err = json.Unmarshal([]byte(data),&r.sessionId)
	if err!=nil{
		return
	}
	return
}

func (r *RedisSession)Del(key string)(err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	r.flag=SessionFlagModify
	delete(r.sessionMap,key)
	return
}
//把session存到redis
func (r *RedisSession) Save()(err error)  {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	//若数据没变，不需要存
	if r.flag!=SessionFlagModify{
		return
	}
	//内存中的session进行序列化
	data, err := json.Marshal(r.sessionMap)
	if err!=nil{
		return
	}
	//获取redis连接
	conn := r.pool.Get()
	//保存kv
	_, err = conn.Do("SET", r.sessionId, string(data))
	//更改状态
	r.flag=SessionFlagNone
	if err!=nil{
		return
	}

	return
}