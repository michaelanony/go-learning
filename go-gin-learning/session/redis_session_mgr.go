package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	address string
	pwd string
	pool *redis.Pool
	rwLock sync.RWMutex
	sessionMap map[string]Session
}

func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	//若有其他参数
	if len(options)>0{
		r.pwd=options[0]
	}
	r.pool=myPool(addr,r.pwd)
	r.address = addr
	return
}
func myPool(addr,password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:      64,
		MaxActive:    100,
		IdleTimeout:  240,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", addr, )
			if err!=nil{
				return nil,err
			}
			//若有密码
			if _,err :=conn.Do("AUTH",password);err!=nil{
				conn.Close()
				return nil,err
			}
			return conn,err
		},
		//连接测试
		//上线注释
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_,err:=c.Do("PING")
			return err
		},
	}
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	//用uuid作为sessionid
	id := uuid.NewV4()
	//转string
	sessionId := id.String()
	//创建session
	session = NewRedisSession(sessionId,r.pool)
	//加入到大map
	r.sessionMap[sessionId] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	session,ok := r.sessionMap[sessionId]
	if !ok{
		err = errors.New("session not exists")
		return
	}
	return
}

func NewRedisSessionMgr() SessionMgr {
	sr:=&RedisSessionMgr{
		address:    "192.168.11.31:30002",
		sessionMap: make(map[string]Session,32),
	}
	return sr
}