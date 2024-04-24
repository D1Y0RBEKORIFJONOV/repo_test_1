package main

import "fmt"

func topMAxTop2Max(nums []int) (int, int) {
	var max1, max2 int

	maxIndex, maxIndex2 := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > max1 {
			max2 = max1
			max1 = nums[i]
			maxIndex2 = maxIndex
			maxIndex = i
		} else if max2 <= nums[i] && nums[i] <= max1 {
			max2 = nums[i]
			maxIndex2 = i
		}
	}

	return maxIndex, maxIndex2
}

func lastStoneWeight(stones []int) int {

	for len(stones) > 1 {
		max1, max2 := topMAxTop2Max(stones)
		if stones[max1] > stones[max2] {
			stones[max1] = stones[max1] - stones[max2]
			stones = append(stones[:max2], stones[max2+1:]...)
		} else {
			stones = append(stones[:max1], stones[max1+1:]...)
			stones = append(stones[:max2-1], stones[max2:]...)

		}
	}

	if len(stones) != 0 {
		return stones[0]
	}
	return 0
}

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}
