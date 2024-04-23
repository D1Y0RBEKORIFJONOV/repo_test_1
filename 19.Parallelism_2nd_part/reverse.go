package main

import "fmt"

/*
Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21

*/

func reverse(x int) int {
	var reverse_x int = 0
	for x != 0 {
		reverse_x = reverse_x*10 + (x % 10)
		x /= 10

		if reverse_x > 2147483647 {
			return 0
		}

		if reverse_x <= -2147483648 {
			return 0
		}
	}

	return reverse_x
}

func main() {
	fmt.Println(reverse(-1234))
}
