package main

func romanToInt(s string) int {
	var result int
	strSlice := []byte(s)
	roman := map[byte]int{
		73: 1,
		86: 5,
		88: 10,
		76: 50,
		67: 100,
		68: 500,
		77: 1000,
	}

	for i, _ := range strSlice {
		if i+1 < len(strSlice) && roman[strSlice[i]] < roman[strSlice[i+1]] {
			result -= roman[strSlice[i]]
		} else {
			result += roman[strSlice[i]]
		}
	}
	return result
}
