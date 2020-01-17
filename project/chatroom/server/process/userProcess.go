package processes

import (
	"encoding/json"
	"go-learning/project/chatroom/common/message"
	"go-learning/project/chatroom/server/utils"
	"net"
)
type UserProcess struct {
	conn net.Conn

}
func (this *UserProcess)serverProcessLogin(conn net.Conn,mes *message.Message) (err error) {
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
	//如果用户id=100，密码等于1，认为合法，否则不合法
	if loginMes.UserId==100 && loginMes.UserPwd=="123456"{
		loginResMes.Code = 200
	}else{
		loginResMes.Code = 400
		loginResMes.Error="该用户不存在"
	}
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
		Conn: this.conn,
	}
	err = tf.writePkg(conn, data)
	return
}