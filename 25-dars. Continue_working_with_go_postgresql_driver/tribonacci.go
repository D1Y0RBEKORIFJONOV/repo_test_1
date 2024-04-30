package main

import "fmt"

func tribonacci(n int) int {
	var (
		x = 0
		z = 1
		y = 1
	)

	for i := 0; i < n; i++ {
		temp := y
		y = x + z + y
		x = z
		z = temp
	}
	return x
}

func main() {
	fmt.Println(tribonacci(4))
}
