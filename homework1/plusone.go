package main

func plusOne(digits []int) []int {
	flag := true
	for i := len(digits) - 1; i >= 0; i-- {
		if flag == true {
			if digits[i]+1 < 10 {

				digits[i] += 1
				flag = false
			} else {
				digits[i] = 0
				if i == 0 {
					digits = append(digits[:0], append([]int{1}, digits[0:]...)...)
					flag = false
				}
			}
		}
	}

	return digits
}
