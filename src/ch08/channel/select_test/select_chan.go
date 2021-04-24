package main

import (
	"fmt"
	"time"
)

/*
	假设要从网上下载一个文件，我启动了 3 个 goroutine 进行下载，
	并把结果发送到 3 个 channel 中。其中，哪个先下载好，就会使用哪个 channel 的结果。
	在这种情况下，如果我们尝试获取第一个 channel 的结果，程序就会被阻塞，
	无法获取剩下两个 channel 的结果，也无法判断哪个先下载好。
	这个时候就需要用到多路复用操作了，
	在 Go 语言中，通过 select 语句可以实现多路复用，其语句格式如下：
	select {
		case i1 = <-c1:
     		//todo
		case c2 <- i2:
			//todo
		default:
			// default todo
	}

	多路复用可以简单理解为:N个channel中,任意一个channel中有数据产生,select都可以监听到
	然后执行相应的分支,接收数据并处理
*/

func main() {
	// 声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	thirdCh := make(chan string)

	// 同时开启3个goroutine 下载
	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		thirdCh <- downloadFile("thirdCh")
	}()

	// 开始select多路复用,哪个channel能获取到值,就说明哪个最先下载好,就用哪个值
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-thirdCh:
		fmt.Println(filePath)
	}

}

func downloadFile(chanName string) string {
	// 模拟下载文件,可以随机time.Sleep试一下
	time.Sleep(time.Second * 2)
	return chanName + ":filePath"
}

/*
	channel为啥是并发安全的?
	channel内部使用了互斥锁来保证并发的安全.
*/
