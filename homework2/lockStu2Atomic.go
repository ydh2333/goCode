package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter2 struct {
	count int32
}

func (c *Counter2) increment3() {
	atomic.AddInt32(&(c.count), 1)
}

func (c *Counter2) GetCounter() int32 {
	return atomic.LoadInt32(&(c.count))
}

func main() {
	var wg sync.WaitGroup

	c := Counter2{}
	for range 10 {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			for range 1000 {
				c.increment3()
			}	
		}()
	}
	wg.Wait()
	
	fmt.Println(c.GetCounter())
}