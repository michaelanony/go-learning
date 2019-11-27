package main

import "fmt"

type Monkey struct{
	Name string
}

func (mk *Monkey) climbing{
	fmt.Println(mk.Name,"Climbing")
}

//LittleMonkey
type LittleMonkey struct{
	Monkey

}
func main() {
	mk = LittleMonkey{
		Monkey{
			Name:"WuKong",
		},
	}
}