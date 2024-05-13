package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	max := max(len(v1), len(v2))
	k := 0
	for i := 0; i < max; i++ {
		k = i
		if i <= len(v1)-1 && i <= len(v2)-1 {

			vi, _ := strconv.Atoi(v1[i])

			v2i, _ := strconv.Atoi(v2[i])

			if vi > v2i {
				return 1
			} else if vi < v2i {
				return -1
			} else {
				continue
			}
		}

		if len(v1) > k {
			for i := k; i < max; i++ {
				vi, _ := strconv.Atoi(v1[i])
				if vi > 0 {
					return 1
				}
			}
		} else if len(v2) > k {

			for i := k; i < max; i++ {
				v2i, _ := strconv.Atoi(v2[i])
				if v2i > 0 {

					return -1
				}
			}

		}
	}
	return 0
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(compareVersion("1.0.0", "2.0.0"))
}
