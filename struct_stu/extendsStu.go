package main

import "fmt"

type A struct {
	a string
}
type B struct {
	A
	a string
	b string
}
type C struct {
	A
	B
	a string
	b string
	c string
}

func main() {
	var c = C{
		a: "c.a",
		b: "c.b",
		c: "c.c",
		
		B: B{
			a: "B.a",
			b: "B.b",
			A: A{
				a:"B.A.a",
			},
		},
		A: A{
			a: "A.a",
		},
	}

	fmt.Println("c.a:", c.a,"\nc.B.a:", c.B.a, "\nc.B.A.a:", c.B.A.a, "\nc.A.a:", c.A.a)
}