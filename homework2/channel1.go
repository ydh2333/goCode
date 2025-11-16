package main

import (
	"fmt"
	"sync"
)

func receive(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("收到：", v)
	}
}

func send(ch chan <- int, wg *sync.WaitGroup){

	defer func ()  {
		close(ch)
		wg.Done()
	}()
	for v := 1; v <= 10; v++ {
		fmt.Println("发送前：", v)
		ch <- v
		fmt.Println("发送后：", v)
	}
}
func main7() {
	var wg sync.WaitGroup
	ch := make(chan int)
	
	wg.Add(1)
	go send(ch, &wg)

	wg.Add(1)
	go receive(ch, &wg)

	wg.Wait()

}