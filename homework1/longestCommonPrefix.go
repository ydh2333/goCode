package main

func longestCommonPrefix(strs []string) string {

	var str string = ""
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 1; i <= len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i>len(strs[j]) || strs[0][:i] != strs[j][:i] {
				return str
			} else {
				if j == len(strs) - 1 {
					str = strs[0][:i]
				}
				
			}
		}
	}

	return str
}
