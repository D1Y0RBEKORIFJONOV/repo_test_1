package main

import "fmt"

func sumNumberic(n int) bool {
	sum := 0

	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum%2 == 0
}

func countEven(num int) int {
	c := 0
	for i := 1; i <= num; i++ {
		if sumNumberic(i) {
			c++
		}
	}
	return c
}

func main() {

	fmt.Println(countEven(26))
}
