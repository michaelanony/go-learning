package main

import "fmt"

func main() {
	age := 19
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
		fallthrough //无条件执行下面的case语句
	case 2,4,5:
		fmt.Println("245")
	default:
		fmt.Println("default")
	}
	switch  {
	case age>2:
		fmt.Println(">2")
	case age==2:
		fmt.Println("2")
	default:
		fmt.Println("<2")
	}

	for i:=0;i<5;i++{
		for j:=0;j<3;j++{
			if i==2 && j==2{
				//break//跳出for循环
				continue
			}
			fmt.Println(i,j)

		}
	}
}
