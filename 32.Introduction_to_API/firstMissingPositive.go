package main

import (
	"fmt"
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	max := nums[0]
	length := len(nums)
	mp := map[int]bool{}
	for i := 0; i < length; i++ {
		mp[nums[i]] = true
	}

	for i := 1; i <= max; i++ {
		if !mp[i] {
			return i
		}
	}
	if max < 0 {
		return 1
	}
	return max + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}
