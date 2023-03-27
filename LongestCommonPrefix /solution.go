package main

func longestCommonPrefix(strs []string) string {
	for index, val := range strs[0] {
		for _, str := range strs[1:] {
			if byte(val) != str[index] || index == len(str) {
				return strs[0][:index]
			}
		}
	}
	return strs[0]
}
