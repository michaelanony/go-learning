package service

import (
	"fmt"
	"go-learning/homeCMS/dao"
	"go-learning/homeCMS/errno"
	"go-learning/homeCMS/model"
)

func RegistrySvc(user *model.HomeUser) (err error) {
	if user.UNickname=="" || user.UPassword==""|| user.UName == "" {
		return errno.ERROR_USER_FORMAT
	}
	nickname := user.UNickname
	_, err = dao.GinDao.GetUser(nickname,"")
	if err==nil{
		return errno.ERROR_USER_EXISTS
	}
	_ ,err = dao.GinDao.RegistryUser(user);
	if err!=nil{
		fmt.Println(err)
	}
	return
}
