package main

import "fmt"

func pointerHomework2(s *[]int) {
	for i := 0; i < len(*s); i++ {
		(*s)[i] *= 2
	}
	for _, v := range *s {
		fmt.Println("homework:", v)
	}
}

func main2() {
	arr1 := []int{1, 2, 3, 4, 5, 6}

	pointerHomework2(&arr1)

	for _, v := range arr1 {
		fmt.Println("main:", v)
	}

}