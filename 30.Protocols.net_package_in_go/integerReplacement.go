package main

import "fmt"

func integerReplacement(n int) int {
	count := 0
	for n != 1 {
		count += 1
		if n%2 == 0 {
			n /= 2
		} else if (n-1)%4 == 0 {
			n -= 1
		} else {
			n++
		}
	}

	return count
}

func main() {
	fmt.Println(integerReplacement(65535))
}
