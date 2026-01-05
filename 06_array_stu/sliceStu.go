package main

import "fmt"

func main1() {
	s1 := []int{10, 20, 30, 40, 50}
	fmt.Println("len:", len(s1), "cap:", cap(s1))
	fmt.Printf("s1原来地址：%p\n", &s1)

	s1 = append(s1, 1)
	fmt.Println(s1)
	fmt.Println("len:", len(s1), "cap:", cap(s1))
	fmt.Printf("s1新地址：%p\n", &s1)

	s1 = append(s1, 2, 3)
	fmt.Println(s1)
	fmt.Println("len:", len(s1), "cap:", cap(s1))

	fmt.Println("---------------------------------")

	// 移除指定元素
	s1 = append(s1[:3], s1[4:]...)
	fmt.Println(s1)
	fmt.Println("len:", len(s1), "cap:", cap(s1))

	fmt.Println("---------------------------------")

	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	//注意：copy不会触发扩容
	copy(dst1, src1)
	copy(dst2, src2)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)
}
