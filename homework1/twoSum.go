package main

import "fmt"

func twoSum(nums []int, target int) []int {
	arr1 := []int{}
outter:
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				arr1 = append(arr1, i, j)
				break outter
			}
		}
	}

	return arr1
}

func main() {
	nums := []int{-1,-2,-3,-4,-5}
	target := -8
	fmt.Println(twoSum(nums, target))
}
