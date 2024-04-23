package main

import (
	"fmt"
	"sync"
)

func addOne(num int, chanNum chan int, wg *sync.WaitGroup) {
	chanNum <- num
	defer wg.Done()
}

func fan_in(result chan int, chan1 chan int, chan2 chan int) {
	for {
		select {
		case num := <-chan1:
			result <- num
		case num := <-chan2:
			result <- num
		default:
			return
		}
	}
}

func main() {

	chan1 := make(chan int, 4)
	chan2 := make(chan int, 4)
	result := make(chan int, 4)
	wait := sync.WaitGroup{}
	wait.Add(4)
	go addOne(1, chan1, &wait)
	go addOne(2, chan1, &wait)
	go addOne(3, chan2, &wait)
	go addOne(4, chan2, &wait)
	wait.Wait()
	fan_in(result, chan1, chan2)
	for {
		select {
		case num := <-result:
			fmt.Println(num)
		default:
			return
		}
	}
}
