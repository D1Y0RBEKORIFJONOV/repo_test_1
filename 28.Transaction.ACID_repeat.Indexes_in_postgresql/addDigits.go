package main

import "fmt"

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	var c = 0
	for c != 1 {
		c = 0
		sum := 0
		for num != 0 {
			sum += num % 10
			num /= 10
			c++
		}
		num = sum
	}
	return num
}
func main() {
	fmt.Println(addDigits(122234))
}
