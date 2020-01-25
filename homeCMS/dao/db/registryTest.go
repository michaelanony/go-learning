package db

import (
	"testing"
)

func init()  {
	//parseTime将mysql中时间类型，自动解释为go结构体中的时间类型
	mysqlDns:="michael:cccbbb@tcp(192.168.11.31:30001)/testDb?parseTime=true"
	err := Init(mysqlDns, "")
	if err!=nil{
		panic(err)
	}
}

//测试获取单个用户信息
func TestGetUser(t *testing.T)  {
	user, err := GetUser("michael")
	if err!=nil{
		panic(err)
	}
	t.Logf("userHandle:%v",user)
}
