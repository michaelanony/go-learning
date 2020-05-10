package main

import (
	"fmt"
	"github.com/gitstliu/go-redis-cluster"
	"time"
)

var (
	UserDao      *MyPool
	RedisCluster *redis.Cluster
)

type MyPool struct {
	BiliCluster *redis.Cluster
}

func NewUserDao(cluster *redis.Cluster) (userDao *MyPool) {
	userDao = &MyPool{
		BiliCluster:   cluster,
	}
	return
}

func createRedisPool() (err error) {
	RedisCluster, err = redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"172.22.33.30:7374", "172.22.33.30:7304"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	return
}

func InitPool() (err error) {
	if err = createRedisPool();err!=nil{
		panic(err)
	}
	UserDao = NewUserDao(RedisCluster)
	return
}
func RedisString(reply interface{}, err error) (string, error) {
	if err != nil {
		return "", err
	}
	switch reply := reply.(type) {
	case []byte:

		return string(reply), nil
	case string:
		return reply, nil
	}
	return "", fmt.Errorf("unexpected type %T for String", reply)
}
func (s *MyPool) GetCodeByCaptchaKey(captchaKey string) (ret *codeRet) {
	command := "sms_" + captchaKey
	res, err := RedisString(s.BiliCluster.Do("GET",command))
	if err != nil {
		fmt.Println("err",err)
		return
	}
	fmt.Println(res)
	//err = json.Unmarshal([]byte(res), ret)
	return

}
