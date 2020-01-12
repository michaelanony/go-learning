package main

import (
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex //全局互斥锁
)

func cal(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func main() {
	for i := 1; i <= 20; i++ {
		go cal(i)
	}
	//sleep 10s
	//time.Sleep(time.Second*10)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
