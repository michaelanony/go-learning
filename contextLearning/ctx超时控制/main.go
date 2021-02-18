package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 传递带超时的上下文告诉阻塞函数
	// 超时过后应该放弃它的工作。
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	ch := make(chan string)
	cb := make(chan bool)
	go t(ch, cb)

Retry:
	for i := 0; i < 3; i++ {
		select {
		case <-ch:
			fmt.Println(2)
		case <-cb:
			fmt.Println(3)
			break Retry
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // 打印"context deadline exceeded"
		}
	}
}
func t(ch chan string, cb chan bool) {
	ch <- "te"
	fmt.Println("sss", len(ch))
	ch <- "te"

	time.Sleep(time.Duration(180 * time.Millisecond))
	cb <- true

}
