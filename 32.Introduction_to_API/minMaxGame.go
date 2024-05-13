package main

import "fmt"

func minMaxGame(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[1]
		}
		return nums[0]
	}
	arr := make([]int, len(nums)/2)
	c := 0
	order := false
	max, min := 0, 0
	for i := 0; i < len(nums)/2; i++ {
		if nums[2*i] < nums[2*i+1] {
			min = nums[2*i]
			max = nums[2*i+1]
		} else {
			max = nums[2*i]
			min = nums[2*i+1]
		}
		if !order {
			order = true
			arr[c] = min
		} else {
			order = false
			arr[c] = max
		}
		c++
	}

	return minMaxGame(arr)
}

func main() {
	fmt.Println(minMaxGame([]int{70, 38, 21, 22}))
}
