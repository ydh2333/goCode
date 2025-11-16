package main

import (
	"fmt"
	"sync"
	"time"
)

func test(s string) {
	startTime := time.Now()
	fmt.Println("执行",s , "任务....")
	time.Sleep(2 * time.Millisecond)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("运行时间：", duration)
}

func main4() {
	var wg sync.WaitGroup
	tasks := []string{"task1", "task2", "task3", "task4"}

	for _, task := range tasks {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			test(s)
		}(task)
	}

	wg.Wait()
}
