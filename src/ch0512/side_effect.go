package main

import "fmt"

func main() {
	n := 0
	result := &n
	Multiply(23, 23, result)
	fmt.Printf("Multiple: %d", *result)
}

func Multiply(a, b int, result *int) {
	*result = a * b
}
