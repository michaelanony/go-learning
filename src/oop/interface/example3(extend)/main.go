package main

import "fmt"

type Monkey struct{
	Name string
}
//声明接口
type BirdAble interface {
	Flying()
}
func (mk *Monkey) climbing(){
	fmt.Println(mk.Name,"Climbing")
}

//LittleMonkey
type LittleMonkey struct{
	Monkey
}
func (mk *LittleMonkey) Flying(){
	fmt.Println(mk.Name,"Flying")
}
func main() {
	mk := LittleMonkey{
		Monkey{
			Name:"WuKong",
		},
	}
	mk.climbing()
	mk.Flying()
}