package main

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-num {
				return []int{i, j}
			}
		}
	}
	panic("error")
}
