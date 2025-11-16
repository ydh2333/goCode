package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	long  float64
	width float64
}

func (r *Rectangle) Area() float64 {
	return r.long * r.width
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.long + r.width)
}

// --------------------------------------------
const PI = 3.14

type Circle struct {
	raduis float64
}

func (c *Circle) Area() float64 {
	return c.raduis * c.raduis * PI
}
func (c *Circle) Perimeter() float64 {
	return 2 * PI * c.raduis
}

func main5() {
	c := Circle{1.0}
	fmt.Println("圆面积：", c.Area())
	fmt.Println("圆周长：", c.Perimeter())

	//---------------------------------------

	r := Rectangle{long: 3, width: 4}
	fmt.Println("矩形面积：", r.Area())
	fmt.Println("矩形周长：", r.Perimeter())

}
