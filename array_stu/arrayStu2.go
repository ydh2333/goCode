package main

import "fmt"

type Custom struct {
	i int
}

var carr [5]*Custom = [5]*Custom{
	{1},
	{2},
	{3},
	{4},
	{5},
}

func main3() {
	a := [5]int{50, 40, 30, 20, 10}
	fmt.Println("before all, a = ", a)
	for i := range carr {
		fmt.Printf("in main func, carr[%d] = %p, value = %v \n", i, carr[i], *carr[i])
		// fmt.Println(carr)
	}
	printFuncParamPointer(carr)
}

func printFuncParamPointer(param [5]*Custom) {
	for i := range param {
		fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, param[i], *param[i])
	}
}
