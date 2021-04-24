package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	现在我们解决了多个 goroutine 同时读写的资源竞争问题，
	但是又遇到另外一个问题——性能。因为每次读写共享资源都要加锁，所以性能低下，
	这该怎么解决呢?
	现在我们分析读写这个特殊场景，有以下几种情况：

		写的时候不能同时读，因为这个时候读取的话可能读到脏数据（不正确的数据）；

		读的时候不能同时写，因为也可能产生不可预料的结果；

		读的时候可以同时读，因为数据不会改变，所以不管多少个 goroutine 读都是并发安全的。

	所以就可以通过读写锁 sync.RWMutex 来优化这段代码，提升性能。
	现在我将以上示例改为读写锁，
*/

var (
	sum     int
	sumSync int
	mutexRW sync.RWMutex
)

func main() {
	// 开启100个协程goroutine 执行sum+10
	for i := 0; i < 1000; i++ {
		//go add(10)
		go addMutex(10)
	}

	for i := 0; i < 10; i++ {
		go readSum()

		go readSumSync()
	}

	//防止提前退出
	time.Sleep(time.Second * 3)
}

/*
	你期待的结果可能是“和为 1000”，但当运行程序后，可能如预期所示，
	但也可能是 990 或者 980。
	导致这种情况的核心原因是资源 sum 不是并发安全的，
	因为同时会有多个协程交叉执行 sum+=i，产生不可预料的结果
*/

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

func readSum() {
	fmt.Println(" 写加锁读不加锁情况: sumSync is : ", sumSync)
}

func readSumSync() {
	mutexRW.RLock()
	defer mutexRW.RUnlock()
	fmt.Println(" 互斥锁优化为读写加锁后的 情况: sumSync is : ", sumSync)

}
