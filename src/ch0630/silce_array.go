package main

import "fmt"

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Printf("b[1:4] is %v\n", b[1:4])
	printSlice(b[1:4])
	fmt.Printf("b[:2] is %v\n", b[:2])
	printSlice(b[:2])
	fmt.Printf("b[2:] is %v\n", b[2:])
	printSlice(b[2:])
	fmt.Printf("b[:] is %v\n", b[:])
	printSlice(b[:])

}

func printSlice(input []byte) {
	for i, v := range input {
		fmt.Printf("index is %d, value is %v, slice value id %v\n", i, v, input[i])
	}
}
