package processes

import (
	"encoding/json"
	"fmt"
	"go-learning/project/chatroom/common/message"
	"go-learning/project/chatroom/server/model"
	"go-learning/project/chatroom/server/utils"
	"net"
)
type UserProcess struct {
	Conn net.Conn

}
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err!=nil{
		panic(err)
		return
	}
	//2、声明一个loginMes完成赋值
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginRes
	//1、使用model.MyUserDao到redis去验证
	user,err:=model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	fmt.Println(user)
	if err !=nil{
		if err == model.ERROR_USER_NOT_EXISTS{
			loginResMes.Code = 500
			loginResMes.Error= err.Error()
		}else if err == model.ERROR_USER_PWD{
			loginResMes.Code = 300
			loginResMes.Error = err.Error()
		}else{
			loginResMes.Code=500
			loginResMes.Error= "服务器内部错误"
		}
	}else{
		loginResMes.Code = 200
		fmt.Println(user,"login success!")
	}
	////如果用户id=100，密码等于1，认为合法，否则不合法
	//if loginMes.UserId==100 && loginMes.UserPwd=="123456"{
	//	loginResMes.Code = 200
	//}else{
	//	loginResMes.Code = 400
	//	loginResMes.Error="该用户不存在"
	//}
	//3、将loginMes序列化
	data,err:=json.Marshal(loginResMes)
	if err!=nil{
		panic(err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5、对resMes进行序列化，准备发送
	data,err =json.Marshal(resMes)
	if err!=nil{
		panic(err)
		return
	}
	//6、发送data，我们将其封装到writePkg
	tf :=&utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data,并直接反序列化成registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&registerMes)
	if err!=nil{
		panic(err)
		return
	}
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err!=nil{
		if err ==model.ERROR_USER_NOT_EXISTS{
			registerResMes.Code=505
			registerResMes.Error=model.ERROR_USER_EXISTS.Error()
		}else{
			registerResMes.Code=506
			registerResMes.Error="Unknown error"
		}
	}else{
		registerResMes.Code=200
	}
	data,err:=json.Marshal(registerResMes.Code)
	if err!=nil{
		panic(err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5、对resMes进行序列化，准备发送
	data,err =json.Marshal(resMes)
	if err!=nil{
		panic(err)
		return
	}
	//6、发送data，我们将其封装到writePkg
	tf :=&utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}