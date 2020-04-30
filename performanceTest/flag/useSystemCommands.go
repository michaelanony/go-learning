package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	flagDemo2()

}

func flagDemo1()  {
	//flag set command's parameters
	//flag.Type(flag's name,default,help infomation)
	//eg: ./flagDemo -name michael -age 12 -marry false
	name:=flag.String("name","michael","姓名")
	age := flag.Int("age",23,"年龄")
	married := flag.Bool("marry",false,"婚否")
	fmt.Println(name,age,married)
}
func flagDemo2()  {
	var(
		name string
		age int
		married bool
		delay time.Duration
	)
	flag.StringVar(&name,"name","michael","姓名")
	flag.IntVar(&age,"age",23,"年龄")
	flag.BoolVar(&married,"marry",false,"婚否")
	flag.DurationVar(&delay,"delay",0,"延迟")
	flag.Parse()//解析命令行参数
	fmt.Println(name,age,married,delay)
}
func normalCommand()  {
	//normal
	if len(os.Args)>0{
		for index,arg :=range os.Args{
			fmt.Printf("args[%d]=%v\n",index,arg)
		}
	}
}