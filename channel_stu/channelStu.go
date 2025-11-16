package main

import (
	"fmt"
	"time"
)

// 只接收channel的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("receiveOnly接收到: %d\n", v)
	}
}

// 只发送channel的函数
func sendOnly(ch chan<- int) {
	defer close(ch)
	for i := range 5 {
		fmt.Printf("发送前: %d\n", i)
		ch <- i
		fmt.Printf("发送后: %d\n", i)
	}

}

func main() {
	// 创建一个带缓冲的channel,可以存储三个数据.数据没有达到缓冲数量之前,可以一直往里面存储数据,满了之后会阻塞
	ch := make(chan int, 3)

	// 启动发送goroutine
	go sendOnly(ch)

	// 启动接收goroutine
	go receiveOnly(ch)

	// 使用select进行多路复用
	// timeout := time.After(2 * time.Second)
	// for {
	// 	select {
	// 	case v, ok := <-ch:
	// 		if !ok {
	// 			fmt.Println("Channel已关闭")
	// 			return
	// 		}
	// 		//time.Sleep(50000 * time.Millisecond)
	// 		fmt.Printf("主goroutine接收到: %d\n", v)
	// 	case <-timeout:
	// 		fmt.Println("操作超时")
	// 		return
	// 	default:
	// 		fmt.Println("没有数据，等待中...")
	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }
}
