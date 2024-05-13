package main

import "fmt"

func majorityElement(nums []int) int {
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
	}
	max := 0
	key := 0
	for k, v := range mp {
		if v > max {
			max = v
			key = k
		}
	}
	return key
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 1}))
}
