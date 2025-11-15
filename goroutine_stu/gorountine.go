package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main2() {
	// 声明等待组，用于让main等待所有goroutine运行完成
	var wg sync.WaitGroup

	// 方式二：
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("run goroutine in closure")
	}()
	// 方式二：
	wg.Add(1)
	go func(s string) {
		defer wg.Done()
		fmt.Println(s)
	}("gorouine: closure params")

	// 方式一：这个方法可以改成以下：
	// go say("in goroutine: world")
    wg.Add(1)
    go func ()  {
        defer wg.Done()
        say("in goroutine: world")
    }()

	say("hello")
	wg.Add(1)
	go func(s string) {
		defer wg.Done()
		for range 5 {
			time.Sleep(60 * time.Millisecond)
			fmt.Println(s)
		}
	}("hhhhhhhhhhhhhhhhhhhhhhhhhhhh")

	// 等待所有goroutine运行完毕
	wg.Wait()
}
