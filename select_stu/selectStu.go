package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch1 <- i
	// 		ch2 <- i
	// 		ch3 <- i
	// 	}
	// }()
	// for i := 0; i < 10; i++ {
	//     select {
	// 	// 读数据
	//     case x := <-ch1:
	//         fmt.Printf("receive %d from channel 1\n", x)
	// 	// 写数据
	//     case ch2 <- 11:
	//         fmt.Printf("send a value to channel 2\n")
	//     case z := <-ch3:
	//         fmt.Printf("receive %d from channel 3\n", z)
	//     }
	// }

	//以下为channel阻塞了之后执行了default
	// for i := 0; i < 10; i++ {
	// 	ch1 <- i
	// 	ch2 <- i
	// 	ch3 <- i
	// }
	// //读取了一个之后就不会阻塞
	// <-ch1
	// for i := 0; i < 1; i++ {
	// 	select {
	// 	case ch1 <- 11:
	// 		fmt.Printf("send a value to channel 1\n")
	// 	case ch2 <- 22:
	// 		fmt.Printf("send a value to channel 2\n")
	// 	case ch3 <- 33:
	// 		fmt.Printf("send a value to channel 3\n")
	// 	default:
	// 		fmt.Println("no channel ready!")
	// 	}
	// }

	// 标准写法:

	for i := range 10 {
		ch1 <- i
		ch2 <- i
		ch3 <- i
	}
	<-ch1
	<-ch2
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	go func() {
		for {
			select {
			case ch1 <- 11:
				fmt.Printf("send a value to channel 1\n")
			case ch2 <- 22:
				fmt.Printf("send a value to channel 2\n")
			case ch3 <- 33:
				fmt.Printf("send a value to channel 3\n")
			case <-ctx.Done():
				fmt.Println("context Done!")
				if err := ctx.Err(); err != nil {
					// todo error info
				}
				return
				// default:
				// 	fmt.Println("no channel ready!")
			}
			fmt.Println("select end!")
		}
	}()
	time.Sleep(4 * time.Second)
}
