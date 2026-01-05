package main

import (
    "fmt"
    "sync"
    "time"
)

// 线程安全的计数器
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

// 增加计数
func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

// 获取当前计数
func (c *SafeCounter) GetCount() int {
    // c.mu.Lock()
    // defer c.mu.Unlock()
    return c.count
}

// -------------------------------------------------------------------
// 线程不安全的计数器
type UnsafeCounter struct {
    count int
}

// 增加计数
func (c *UnsafeCounter) Increment() {
    c.count += 1
}

// 获取当前计数
func (c *UnsafeCounter) GetCount() int {
    return c.count
}

func main() {
    counter := SafeCounter{}

    // 启动1000个goroutine同时增加计数
    for range 1000 {
        go func() {
            for range 100 {
                counter.Increment()
            }
        }()
    }

    // 等待一段时间确保所有goroutine完成
    time.Sleep(time.Second)

    // 输出最终计数
    fmt.Printf("Final count: %d\n", counter.GetCount())
}