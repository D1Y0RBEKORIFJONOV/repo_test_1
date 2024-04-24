package main

import "fmt"

/*
ou are given an integer array nums of even length. You have to split the array into two parts nums1 and nums2 such that:

nums1.length == nums2.length == nums.length / 2.
nums1 should contain distinct elements.
nums2 should also contain distinct elements.
Return true if it is possible to split the array, and false otherwise.



Example 1:

Input: nums = [1,1,2,2,3,4]
Output: true
Explanation: One of the possible ways to split nums is nums1 = [1,2,3] and nums2 = [1,2,4].
Example 2:

Input: nums = [1,1,1,1]
Output: false
Explanation: The only possible way to split nums is nums1 = [1,1] and nums2 = [1,1]. Both nums1 and nums2 do not contain distinct elements. Therefore, we return false.
*/

func isPossibleToSplit(nums []int) bool {
	nums1 := map[int]int{}
	nums2 := map[int]int{}

	for _, num := range nums {
		_, ok := nums2[num]
		if !ok {
			nums2[num] = 1
		} else {
			nums1[num] = 1
		}
	}
	if (len(nums1) + len(nums2)) != len(nums) {
		return false
	}
	return true
}

func main() {
	fmt.Println(isPossibleToSplit([]int{1, 1, 2, 2}))
	/*  6 3 1 9

	1 1 8 2
	*/
}
