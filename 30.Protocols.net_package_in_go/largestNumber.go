package main

import (
	"fmt"
	"strconv"
)

func countX(n int) int {
	c := 0
	for n != 0 {
		n /= 10
		c++
	}
	return c
}
func finMax(nums []int) int {
	m := nums[0] % 10
	x := countX(nums[0])
	index := 0

	for i := 0; i < len(nums); i++ {
		if m < nums[i]%10 && x >= countX(nums[i]) {
			m = nums[i] % 10
			x = countX(nums[i])
			index = i
		}
	}

	return index
}

func largestNumber(nums []int) string {
	str := ""
	for len(nums) > 1 {
		index := finMax(nums)
		str += strconv.Itoa(nums[index])
		nums = append(nums[:index], nums[index+1:]...)
	}
	str += strconv.Itoa(nums[0])
	return str
}
func main() {

	fmt.Println(largestNumber([]int{34323, 3432}))

}
