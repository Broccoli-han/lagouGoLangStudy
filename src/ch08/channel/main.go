package main

import "fmt"

func main() {
	/*
		使用make创建的chan 是一个无缓冲的channel，
		其容量为0，不能存储任何数据
		无缓冲的channel只能起到传输数据的作用，数据不会在channel做任何停留
		这也意味着，无缓冲的channel的发送和接收操作是同时进行的，它亦可称为同步channel
	*/
	ch := make(chan string)

	go func() {
		fmt.Println("broccoli")
		ch <- "goroutine finished"
	}()

	fmt.Println("i am main goroutine")

	v := <-ch
	fmt.Println("receive the value in chan is :", v)

	/*
		运行此示例 发现程序并没有退出，达到了time.Sleep()函数的效果
		在新启动的goroutine中 向chan类型的变量ch中发送值，在 main gouroutine
		中 从变量ch中接收值，如果ch中没有值，则阻塞等待到ch中的值可以被接收为止
	*/
}
