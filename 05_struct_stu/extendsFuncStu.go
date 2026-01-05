package main

import (
	"fmt"
)

type A struct {
	a string
}

func (a A) GetA() string {
	return a.a
}

func (a A) SetA(s string) {
	a.a = s
	fmt.Printf("a的内存地址：%p\n", &a)
}

func (a *A) SetPA(s string) {
	a.a = s
	fmt.Printf("a2的内存地址：%p\n", a)

}

func (a *A) GetPA() string {
	return a.a
}

type B struct {
	A
	b string
	c []int
	d int
}

func (b *B) modifyB() {
	b.b = "modify B.b"
}

func (b B) modifyC() {
	b.c[1] = 100
}

func (b *B) modifyD() {
	b.d = 200
}

func NewB() B {
	return B{
		b: "String Bb",
		c: []int{1, 2, 3},
		d: 1,
		A: A{
			a: "String Aa",
		},
	}
}

func main1() {
	var a A
	a.SetA("set a!")

	fmt.Printf("a的内存地址：%p\n", &a)

	resa := a.GetA()
	fmt.Println("a.GetA():", resa)

	fmt.Println("--------------------")

	var a2 A

	a2.SetPA("set PA!")

	fmt.Printf("a2的内存地址：%p\n", &a2)

	respa := a2.GetPA()
	fmt.Println(respa)

	fmt.Println("--------------------")

	var b B
	b = NewB()

	fmt.Println(b)

	b.modifyC()
	b.modifyD()
	b.modifyB()

	fmt.Println(b)

}
