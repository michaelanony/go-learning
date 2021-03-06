package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/project/chatroom/client/utils"
	"go-learning/project/chatroom/common/message"
	"net"
	"os"
)

type UserProcess struct {

}

func (this *UserProcess)Login(userId int,userPwd string)(err error)  {
	//fmt.Printf("userId=%d userPwd=%s\n",userId,userPwd)
	//1.连接到服务器
	conn,err:=net.Dial("tcp","127.0.0.1:8880")
	if err!=nil{
		panic(err)
		return
	}
	//延时关闭
	defer conn.Close()
	//2、准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//3、创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd=userPwd
	//4、将loginMes序列化
	data,err:=json.Marshal(loginMes)
	if err!=nil{
		panic(err)
		return
	}
	//5、把data赋值给mes.Data字段
	mes.Data = string(data)
	//6、将mes进行序列化
	data,err =json.Marshal(mes)
	if err!=nil{
		panic(err)
		return
	}
	//7、到这个时候，data就是我们要的消息
	//7.1 先把data长度发送给服务器（防止丢包）
	//先获取到data到长度->转成一个表示长度到byte切片（conn.Write获取的是byte[]）
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)//把长度转换为了bytes
	//发送长度
	_, err =conn.Write(buf[:4])
	if err!=nil{
		fmt.Println("conn.write failed",err)
		return
	}

	//fmt.Printf("Client sent ok,length is %d,content is %s",len(data),string(data))
	_,err = conn.Write(data)
	if err !=nil{
		panic(err)
		return
	}

	//这里处理服务端返回到消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err!=nil{
		panic(err)
		return
	}
	//将mes到data部分反序列化成LoginResMes
	var loginResMes message.LoginRes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code ==200{
		fmt.Println("Login Success")
		//启动一个协程保持和服务端通讯，如果服务端有数据推送给客户端
		go serverProcessMes(conn)
		//1、调用登入成功后的菜单
		for{
			ShowMenu()
		}
	}else{
		fmt.Println(loginResMes.Error)
	}
	return
}

func (this *UserProcess)Register(userId int,userPwd string,userName string)(err error)  {
	//1.连接到服务器
	conn,err:=net.Dial("tcp","127.0.0.1:8880")
	if err!=nil{
		panic(err)
		return
	}
	//延时关闭
	defer conn.Close()
	//2、准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3、创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd=userPwd
	registerMes.User.UserName=userName
	//4、将registerMes序列化
	data,err:=json.Marshal(registerMes)
	if err!=nil{
		panic(err)
		return
	}
	//5、把data赋值给mes.Data字段
	mes.Data = string(data)
	//6、将mes进行序列化
	data,err =json.Marshal(mes)
	if err!=nil{
		panic(err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err=tf.WritePkg(data)
	if err!=nil{
		panic(err)
		return
	}
	mes,err = tf.ReadPkg()
	if err!=nil{
		panic(err)
		return
	}
	//将mes到data部分反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code ==200{
		fmt.Println("注册成功")
	}else{
		fmt.Println(registerResMes.Error)
	}
	os.Exit(0)
	return
}
