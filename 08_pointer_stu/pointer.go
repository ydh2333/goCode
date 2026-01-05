package main

import "fmt"

func main() {
	var p1 *int
	var p2 *string

	i := 1
	s := "hello"

	p1 = &i
	p2 = &s

	fmt.Println("*p1:", *p1, ", &i:", &i, "p1 address:", p1)
	fmt.Println("*p2:", *p2, ", &s:", &s, "p2 address:", p2)

	fmt.Println("-----------------------")

	*p1 = 10
	fmt.Println("After Modify *p1:", *p1, ", i:", i, "p1 address:", p1)

	fmt.Println("-----------------------")

	a := 2
	var p *int
	fmt.Println(&a)
	p = &a
	fmt.Println(p, &a)

	var pp **int
	pp = &p
	fmt.Println(pp, p)
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p)
	fmt.Println(a, &a)

	fmt.Println("-----------------------")

}
