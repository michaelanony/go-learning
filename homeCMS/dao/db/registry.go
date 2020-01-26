package db

import (
	"go-learning/homeCMS/model"
)

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
func (d *UserDao)GetUser(nickName string) (ret *model.HomeUser,err error) {
	ret = &model.HomeUser{}
	sqlStr:="select * from home_user where u_nickname = ?"
	err = d.MysqlPool.Get(ret,sqlStr,nickName)
	return
}

