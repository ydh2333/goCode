package main

import (
	"fmt"
	"sync"
	"time"
)

func main3() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				time.Sleep(80 * time.Millisecond)
				fmt.Println("go1偶数:", i)
			}
		}
	}()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				time.Sleep(40 * time.Millisecond)
				fmt.Println("go2奇数:", i)
			}
		}
	}()

	wg.Wait()

}
