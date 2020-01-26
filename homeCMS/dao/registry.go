package dao

import (
	"github.com/garyburd/redigo/redis"
	"go-learning/homeCMS/model"
)
//MYSQL PROCESS

//注册用户
func (d *UserDao)RegistryUser(homeUser *model.HomeUser) (userId int64 ,err error) {
	sqlStr:="insert into home_user(u_name,u_nickname,u_password,u_register_ip) values(?,?,?,?)"
	ret, err := d.MysqlPool.Exec(sqlStr, homeUser.UName, homeUser.UNickname, homeUser.UPassword,homeUser.URegisterIp)
	if err!=nil{
		return
	}
	userId,err=ret.LastInsertId()
	return
}

//从数据库获取用户
func (d *UserDao)GetUser(nickName,password string) (ret *model.HomeUser,err error) {
	ret = &model.HomeUser{}
	if password ==""{
		sqlStr:="select * from home_user where u_nickname = ?"
		err = d.MysqlPool.Get(ret,sqlStr,nickName)
		return
	}
	sqlStr :="select * from home_user where u_nickname = ? and u_password=?"
	err = d.MysqlPool.Get(ret,sqlStr,nickName,password)
	return
}
//检查用户名密码是否正确


//REDIS PROCESS
//往redis插入已注册IP防止频繁注册
func (d *UserDao)InsetIntoRedis(ip string)(err error){
	conn := d.RedisPool.Get()
	if _,err = redis.String(conn.Do("SETEX",ip,10,1));err!=nil{
		return
	}
	return
}

//查询redis中是否存在该IP
func (d *UserDao)CheckIpInRedis(ip string)(flag bool,err error)  {
	flag = true
	conn := d.RedisPool.Get()
	ret, err := redis.String(conn.Do("GET", ip))
	if ret==""{
		flag = false
		return
	}
	return
}