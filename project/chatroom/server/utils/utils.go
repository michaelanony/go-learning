package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/project/chatroom/common/message"
	"net"
)
//我们将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte //传输时使用的缓冲

}
func (this *Transfer) ReadPkg() (mes message.Message,err error) {
	fmt.Println("Reading data....")

	_,err =this.Conn.Read(this.Buf[:4])
	if err!=nil{
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	n,err:=this.Conn.Read(this.Buf[:pkgLen])
	if n!=int(pkgLen) ||err!=nil{
		return
	}

	//反序列化（把pkglen反序列化成message.Message)
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err !=nil{
		panic(err)
		return
	}
	return

}
func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)//把长度转换为了bytes
	//发送长度
	n, err :=this.Conn.Write(this.Buf[:4])
	if err!=nil || n!=4{
		fmt.Println("conn.write failed",err)
		return
	}
	//发送data本身
	n, err =this.Conn.Write(data)
	if err!=nil || n!=int(pkgLen){
		fmt.Println("conn.write failed",err)
		return
	}
	return
}