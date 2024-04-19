package main

import "fmt"

func GetFactorial(num int, sum chan int) {
	if num == 0 {
		close(sum)
		return
	}
	sum <- num
	go GetFactorial(num-1, sum)
}

func main() {
	var num int
	fmt.Print("Son kirting: ")
	fmt.Scan(&num)
	ch := make(chan int)
	go GetFactorial(num, ch)
	sum := 1
	for v := range ch {
		sum *= v
	}
	fmt.Println(sum)

}
