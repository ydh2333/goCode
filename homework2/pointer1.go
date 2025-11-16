package main

import "fmt"

func pointerHomework(p *int) int {
	*p += 10
	return *p
}

func main1() {
	p := 0
	fmt.Println("函数输出的值：", pointerHomework(&p))
	fmt.Println("变量本身的值：", p)
}
