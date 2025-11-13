package main

func appearOne(arr []int) int {
	map1 := map[int]int{}

	for i := 0; i < len(arr); i++ {
		map1[arr[i]] += 1
	}

	for k, v := range map1 {
		if v == 1 {
			return k
		}
	}

	return -1
}
