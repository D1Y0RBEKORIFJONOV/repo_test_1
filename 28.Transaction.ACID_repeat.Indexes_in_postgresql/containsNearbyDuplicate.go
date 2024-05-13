package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	for i, v := range nums {
		if _, ok := mp[v]; ok {
			if i-mp[v] <= k {
				return true
			}
		}
		mp[v] = i
	}
	return false
}
func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 4, 5, 6, 7}, 5))
}
