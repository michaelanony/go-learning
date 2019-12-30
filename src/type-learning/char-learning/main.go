package main

import "fmt"

func main(){
	s1:="Golang"
	c1:='G'//ASCLL码下占一个字节（8位）
	fmt.Println(s1,c1)
	s2:="中国"
	c2:='中' //UTF-8编码一个中文占3位
	fmt.Println(s2,c2)
	s3:="hello中国"
	fmt.Println(len(s3))
	for i:=0;i<len(s3);i++{
		fmt.Printf("%c\n",s3[i])
	}
	fmt.Println()
	//for range 按照rune类型去遍历
	for _,v:=range s3{
		fmt.Printf("%c\n",v)
	}
	s5:="big"
	byterArray :=[]byte(s5)
	fmt.Println(byterArray)
	byterArray[0] = 'p'
	s5 = string(byterArray)
	fmt.Println(s5)

	//作业 字符串倒序
	s6 := "hello"
	tmpStringArray := []byte(s6)
	for t:=0 ;t<len(tmpStringArray)/2;t++{
		tmpStringArray[t],tmpStringArray[len(tmpStringArray)-1-t]= tmpStringArray[len(tmpStringArray)-1-t],tmpStringArray[t]
	}
	s6 = string(tmpStringArray)
	fmt.Println(s6)

}

