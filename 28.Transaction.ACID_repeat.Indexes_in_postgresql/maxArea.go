package main

import "fmt"

func maxArea(height []int) int {
	a := 0
	b := len(height) - 1
	aMax := 0
	for a < b {
		ar := min(height[a], height[b]) * (b - a)
		if ar > aMax {
			aMax = ar
		}
		if height[a] < height[b] {
			a++
		} else {
			b--
		}
	}
	return aMax
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxArea([]int{1, 2, 3, 4}))
}
