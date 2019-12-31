package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 077
	var c int = 0xff
	fmt.Println(a, b, c)
	fmt.Printf("%b\n", a)
	fmt.Printf("%o\n", b)
	fmt.Printf("%x\n", c)
	x := 1
	p := &x
	p2 := &p
	fmt.Println(p, p2, *p2, **p2)
	*p = 2
	fmt.Println(x, p2)
	pi := new(int)
	fmt.Println(*pi)

}
