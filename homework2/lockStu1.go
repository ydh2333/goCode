package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mx     sync.Mutex
	counte int
}

func (c *Counter) Increment2() {
	c.mx.Lock()
	defer c.mx.Unlock()
	time.Sleep(1 * time.Millisecond)
	c.counte++
}
func (c *Counter) GetCount2() int {
	return c.counte
}

func main9() {
	var wg sync.WaitGroup
	c := Counter{}

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				c.Increment2()
			}
		}()
	}
	wg.Wait()
	fmt.Println(c.GetCount2())

}
