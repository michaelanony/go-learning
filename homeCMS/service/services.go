package service

import (
	"fmt"
	"go-learning/homeCMS/dao/db"
	"go-learning/homeCMS/errno"
	"go-learning/homeCMS/model"
)

func RegistrySvc(user *model.HomeUser) (err error) {
	if user.UNickname=="" || user.UPassword==""|| user.UName == "" {
		return errno.ERROR_USER_FORMAT
	}
	nickname := user.UNickname
	_, err = db.GetUser(nickname)
	if err==nil{
		return errno.ERROR_USER_EXISTS
	}
	userId ,err := db.RegistryUser(user);
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("User id is %v",userId)

	return
}
