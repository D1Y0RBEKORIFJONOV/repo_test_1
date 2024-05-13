package main

import (
	"fmt"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	sum := 0

	for i := 0; i < k; i++ {
		if happiness[i]-i >= 0 {
			sum += happiness[i] - i
		}
	}
	return int64(sum)
}

func main() {
	fmt.Print(maximumHappinessSum([]int{2, 83, 62}, 3), "\n")

}
