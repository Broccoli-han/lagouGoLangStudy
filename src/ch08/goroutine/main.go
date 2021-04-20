package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("broccoli")
	fmt.Println("我是 main goroutine")
	/*此处让 main goroutine 等一秒，
	  不然main goroutine执行完毕，程序就退出了，
	   也就看不到新的goroutine 中打印的语句了
	*/
	// 此处让main goroutine 等一秒
	time.Sleep(time.Second * 1)

}
