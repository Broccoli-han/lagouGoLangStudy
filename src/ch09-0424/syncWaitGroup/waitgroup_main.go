package main

import (
	"fmt"
	"sync"
)

/*
	多协程时,需要监听所有协程都执行完毕,才能执行main函数退出-结束程序
	需要用到sync.WaitGroup
	使用方法分三步:
		1> 声明一个sync.WaitGroup ,然后通过Add方法设置计数器的值,根据需要跟踪的协程数来设置
		2> 在每个协程执行完毕时,调用Done方法,计数器减1,告诉sync.WaitGroup该协程已经执行完毕
		3> 最后调用Wait方法一直等待,直到计数器值变为0,即所有跟踪的协程都执行完毕.
*/

var (
	mutexRW sync.RWMutex
	sumSync int
)

func main() {
	run()
}

func run() {
	var wg sync.WaitGroup
	// 因为需要监视1010个协程,所以设置计数器为1010
	wg.Add(1100)

	// 开启1000个协程goroutine 执行sum+10
	for i := 0; i < 1000; i++ {

		go func() {
			defer wg.Done()
			addMutex(10)
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			readSumSync()
		}()
	}

	// 一直等待,只要计数器值为0
	wg.Wait()
}

func addMutex(i int) {
	/*
		Mutex中的Lock\Unlock总是成对出现,而且要确保Lock获得锁后,一定执行Unlock
		所以在函数或者方法中 使用defer 语句释放锁
	*/
	mutexRW.Lock()
	defer mutexRW.Unlock()

	sumSync += i
}

/*
	对sum的加法进行了加锁,可以保证在修改sum值时,并发安全,
	但是,如果读取操作也是采用多协程的呢?
*/

func readSumSync() {
	mutexRW.RLock()
	defer mutexRW.RUnlock()
	fmt.Println(" 互斥锁优化为读写加锁后的 情况: sumSync is : ", sumSync)

}
