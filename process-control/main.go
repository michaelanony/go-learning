package main

import "fmt"

func main() {
	age := 26
	if age>18{
		fmt.Println("over 18")
	}else if age<18{
		fmt.Println("under 18")
	}else{
		fmt.Println("equal 18")
	}

	if age2:=28;age2>18{
		fmt.Println("over 18")
	}
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	for age>0{
		fmt.Println(age)
		age = age-1
	}
	switch age {
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("default")
	}
	age2 :=26
	switch  {
	case age2<18:
		fmt.Println("<18")
		fallthrough //无条件执行下一个case语句
	case age2 >18&&age2<25:
		fmt.Println(">18<25")
	case age<30:
		fmt.Println("<30")
	default:
		fmt.Println(">30")
	}
	for i:=0;i<5;i++{
		for j:=0;j<3;j++{
			if i==2 && j==2{
				break
			}
		}
	}

}
