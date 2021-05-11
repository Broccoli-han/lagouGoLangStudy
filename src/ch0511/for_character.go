package main

import "fmt"

func main() {
	//1 使用两层for循环
	fmt.Printf("use 2 nested for loops\n")
	for i := 0; i < 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Printf("\n")
	}

	// 2 仅用一层for循环及字符串连接
	fmt.Printf("use only one for loop and string concatenation\n")
	str := "G"
	for i := 0; i < 25; i++ {
		println(str)
		str += "G"
	}
}
