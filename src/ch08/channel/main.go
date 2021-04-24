package channel

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
		fmt.Println("new goroutine is print: broccoli")
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

	/*
		有缓冲的通道channel类似一个队列,可以获取其容量和里面元素的个数
	*/
	cacheCh := make(chan int, 5)
	cacheCh <- 2
	cacheCh <- 3
	cacheCh <- 4
	fmt.Printf("cacheCh容量为: %v,元素个数为:%v", cap(cacheCh), len(cacheCh))

	// 通道还可以使用内置函数 close()关闭,通道关闭,不能向其中发送元素,但是可以从其中接收数据,
	// 若channel 中没有数据的话, 接受的数据是元素类型的默认值

	/*
		单向channel, 因业务需求,要限制一个channel只可以接收不能发送,或者只可以发送不能接收,
		这种channel称为单项channel:
		声明方式也很简单,只需带上<-操作符即可
	*/
	//onlySendCh := make(chan<- int)
	//onlyReceiveCh := make(<-chan int)

}

/*
	在函数或方法的参数中,使用单向channel的比较多,这样可以防止一些操作影响了channel
*/
func counter(out chan<- int) {
	// 函数内容使用变量out,只能进行发送操作
}
