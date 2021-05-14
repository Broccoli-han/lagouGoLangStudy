package main

import "fmt"

func main() {
	//fmt.Printf("a is %d , b is %d, and  sum , sub , multi is %v",2,9,calcInt(2,9))
	fmt.Println(calcInt(2, 9))
	fmt.Println(calcInt2(3, 9))
}

func calcInt(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}
func calcInt2(a, b int) (sum int, sub int, multi int) {
	//sum = a + b
	//sub = a - b
	//multi = a * b
	sum, sub, multi = a+b, a-b, a*b
	return
}
