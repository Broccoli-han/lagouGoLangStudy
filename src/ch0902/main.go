package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 0, 1, 0, 1, 0, 1}
	fmt.Println(DeleteSlice2(arr))
}

func DeleteSlice2(a []int) []int {
	j := 0
	for _, val := range a {
		if val == 1 {
			a[j] = val
			j++
		}
	}
	return a[:j]
}
