package main

import (
	"fmt"
)

// 转字符串：
// func huiwen(num int) bool {
// 	str1 := strconv.Itoa(num)

// 	for i, j := 0, int(len(str1)-1); i < j; i++ {
// 		fmt.Println("第", i, "趟", j, ":", str1[i], str1[j])
// 		if str1[i] != str1[j] {
// 			return false
// 		}
// 		j--
// 	}
// 	return true
// }

// 不转字符串
func huiwen(num int) bool {
	if num < 0 {
		return false
	}

	var num2 int = num
	var num3 int = 0
	for ; num2 != 0; num2 /= 10 {
		num3 = num2%10 + num3*10
	}
	fmt.Println(num, num3)
	if num != num3 {
		return false
	}

	return true
}