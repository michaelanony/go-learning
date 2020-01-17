package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/project/chatroom/common/message"
	"net"
)

func login(userId int,userPwd string)(err error)  {
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
	mes, err = readPkg(conn)
	if err!=nil{
		panic(err)
		return
	}
	//将mes到data部分反序列化成LoginResMes
	var loginResMes message.LoginRes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code ==200{
		fmt.Println("Login Success")
	}else{
		fmt.Println(loginResMes.Error)
	}
	return
}
