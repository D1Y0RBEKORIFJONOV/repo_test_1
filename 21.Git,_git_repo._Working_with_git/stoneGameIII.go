package main

import "fmt"

func stoneGameIII(stones []int) string {
	n := len(stones)
	m := 3

	dp := make([]int, m+1)
	sum := make([]int, m+1)

	for i := 0; i < m; i++ {
		stones = append(stones, 0)
	}
	for i := n - 1; i >= 0; i-- {
		for j := m; j > 0; j-- {
			sum[j] += stones[i] - stones[i+m-j]
			dp[j] = dp[j-1]
		}
		sum[0] += stones[i] - stones[i+m]

		dp[0] = sum[0] - dp[m]
		for j := 1; j < m; j++ {
			if sum[j]-dp[m-j] > dp[0] {
				dp[0] = sum[j] - dp[m-j]
			}
		}
	}
	if dp[0] == 0 {
		return "Tie"
	}
	if dp[0] > 0 {
		return "Alice"
	}
	return "Bob"
}

func main() {
	fmt.Println(stoneGameIII([]int{-1, -2, -3}))

}
