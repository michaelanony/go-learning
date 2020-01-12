package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	var intChan chan int
	//创建可以存放3个int类型的管道
	intChan = make(chan int, 3)
	//看看intchan是什么
	fmt.Printf("%v,%v", intChan, &intChan)
	//向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	//看看管道长度和容量
	fmt.Println("channel len =%v cap=%v\n", len(intChan), cap(intChan))
	//取数据
	num1 := <-intChan
	num2 := <-intChan
	fmt.Println(num1, num2)

	//创建接受所有类型的信道
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	cat := Cat{Name: "tom", Age: 18}
	allChan <- cat
	allChan <- 10
	allChan <- "m"
	newCat := <-allChan
	a := newCat.(Cat) //必须使用断言
	fmt.Printf("newCat.Name=%v", a.Name)

	//channel关闭后就不能往里面写数据，但是可以从该channel读取数据
	intChan2 := make(chan int, 3)
	intChan2 <- 100
	intChan2 <- 200
	close(intChan2)
	fmt.Println("close channel")
	n1 := <-intChan2
	fmt.Println(n1)
	//遍历管道
	intChan3 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan3 <- i * 2
	}
	t1 := <-intChan3
	fmt.Println(t1)
	//必须关闭管道才能遍历
	close(intChan3)
	for v := range intChan3 {
		fmt.Println("v=", v)
	}
}
