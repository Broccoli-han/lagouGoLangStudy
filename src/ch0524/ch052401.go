package main

import (
	"fmt"
)

func transferRegionFromWebToData(regions []int) int {
	region := 0
	var arr []int
	fmt.Println(len(arr))
	for _, value := range regions {
		if value != 0 {
			arr = append(arr, value)
		}
	}
	length := len(arr)
	if length == 0 {
		return region
	}
	if length > 1 {
		region = 3
	} else {
		region = arr[0]
	}

	fmt.Println(region)
	return region
}

func main() {
	transferRegionFromWebToData([]int{0, 1})
}
