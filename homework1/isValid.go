package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	sSlice := []rune(s)

	arr1 := make([]rune, 0)

	for i := 0; i <= len(s)-1; i++ {
		if len(arr1) == 0 {
			if sSlice[i] == ')' || sSlice[i] == '}' || sSlice[i] == ']' {
				return false
			}
		}
		// fmt.Println(arr1)
		if sSlice[i] == '(' || sSlice[i] == '{' || sSlice[i] == '[' {
			arr1 = append(arr1, sSlice[i])
		}

		if sSlice[i] == ')' {
			if arr1[len(arr1)-1] == '(' {
				arr1 = append(arr1[:len(arr1)-1])
			} else {
				return false
			}
		}
		if sSlice[i] == ']' {
			if arr1[len(arr1)-1] == '[' {
				arr1 = append(arr1[:len(arr1)-1])
			} else {
				return false
			}
		}
		if sSlice[i] == '}' {
			if arr1[len(arr1)-1] == '{' {
				arr1 = append(arr1[:len(arr1)-1])
			} else {
				return false
			}
		}
	}
	if len(arr1) != 0 {
		return false
	}

	return true
}
