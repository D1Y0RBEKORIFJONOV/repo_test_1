package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isLeapYear(year int) bool {

	if year%4 != 0 {
		return false
	}

	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}

func dayOfYear(date string) int {
	st := strings.Split(date, "-")
	days := 0
	year, _ := strconv.Atoi(st[0])
	month, _ := strconv.Atoi(st[1])
	day, _ := strconv.Atoi(st[2])
	n := 30
	n1 := 31
	for i := 1; i < month; i++ {
		if i == 2 && isLeapYear(year) {
			days += 29
		} else if i == 2 {
			days += 28
		} else if i%2 != 0 {
			days += n1
		} else if i == 8 {
			days += n1
			n, n1 = n1, n
		} else {
			days += n
		}
	}
	return days + day
}

func main() {

	fmt.Println(dayOfYear("2000-12-04"))
}
