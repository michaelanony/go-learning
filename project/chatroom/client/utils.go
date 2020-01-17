package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/project/chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message,err error) {
	fmt.Println("Reading data....")
	buf := make([]byte,8096)
	_,err =conn.Read(buf[:4])
	if err!=nil{
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n,err:=conn.Read(buf[:pkgLen])
	if n!=int(pkgLen) ||err!=nil{
		return
	}

	//反序列化（把pkglen反序列化成message.Message)
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err !=nil{
		panic(err)
		return
	}
	return

}
func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)//把长度转换为了bytes
	//发送长度
	n, err :=conn.Write(buf[:4])
	if err!=nil || n!=4{
		fmt.Println("conn.write failed",err)
		return
	}
	//发送data本身
	n, err =conn.Write(data)
	if err!=nil || n!=int(pkgLen){
		fmt.Println("conn.write failed",err)
		return
	}
	return
}