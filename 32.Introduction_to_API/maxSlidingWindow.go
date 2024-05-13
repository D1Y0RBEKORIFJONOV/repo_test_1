package main

import "fmt"

func findMax(nums []int, startIndex, endIndex int) int {
	max := 0
	fmt.Println(nums, startIndex, endIndex)
	for i := startIndex; i <= endIndex; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	for i := 0; i <= len(nums)-k; i++ {
		result = append(result, findMax(nums, i, k+i))
	}
	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, -1, 45, 34}, 1))
}
