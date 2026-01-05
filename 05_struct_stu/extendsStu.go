package main

import "fmt"

type A1 struct {
	a string
}
type B1 struct {
	A1
	a string
	b string
}
type C1 struct {
	A1
	B1
	a string
	b string
	c string
}

func main3() {
	var c = C1{
		a: "c.a",
		b: "c.b",
		c: "c.c",
		
		B1: B1{
			a: "B.a",
			b: "B.b",
			A1: A1{
				a:"B.A.a",
			},
		},
		A1: A1{
			a: "A.a",
		},
	}

	fmt.Println("c.a:", c.a,"\nc.B.a:", c.B1.a, "\nc.B.A.a:", c.B1.A1.a, "\nc.A.a:", c.A1.a)
}