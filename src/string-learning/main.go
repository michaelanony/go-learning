package main

import ("fmt"
"strings"
)

func main(){
	//拼接字符串
	s1:="alex"
	s2:="eat"
	fmt.Println(len(s1))
	fmt.Println(s1+s2)
	s3 := fmt.Sprintf("%s---%s",s1,s2)
	fmt.Println(s3)
	//分割字符串
	ret := strings.Split(s1,"x")
	fmt.Println(ret)
	//判断是否包含
	ret2 := strings.Contains(s1,"ea")
	fmt.Println(ret2)
	//判断前缀和后缀
	ret3 := strings.HasPrefix(s1,"a")
	fmt.Println(ret3)
	ret4 := strings.HasSuffix(s1,"x")
	fmt.Println(ret4)
}
