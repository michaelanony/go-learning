package db

import (
	"go-learning/homeCMS/model"
)

//注册用户
func RegistryUser(homeUser *model.HomeUser) (userId int64 ,err error) {
	sqlStr:="insert into home_user(u_name,u_nickname,u_password)"
	ret, err := MysqlDB.Exec(sqlStr, homeUser.UName, homeUser.UNickname, homeUser.UPassword)
	if err!=nil{
		return
	}
	userId,err=ret.LastInsertId()
	return
}

//从数据库获取用户
func GetUser(nickName string) (ret *model.HomeUser,err error) {
	ret = &model.HomeUser{}
	sqlStr:="select * from home_user where u_nickname = ?"
	err = MysqlDB.Get(ret,sqlStr,nickName)
	return
}