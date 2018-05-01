package main

import (
	"fmt"
)

func hasPairWithSum(a []int, sum int) bool {
	i, j := 0, len(a)-1
	for i < j {
		s := a[i] + a[j]
		if s > sum {
			j--
		} else if s < sum {
			i++
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(hasPairWithSum([]int{1, 2, 3, 9}, 8))
	fmt.Println(hasPairWithSum([]int{1, 2, 4, 4}, 8))
}
