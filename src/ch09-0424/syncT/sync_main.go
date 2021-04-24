package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享资源
var (
	sum     int
	sumSync int
	mutex   sync.Mutex
)

func main() {
	// 开启100个协程goroutine 执行sum+10
	for i := 0; i < 1000; i++ {
		go add(10)
		go addMutex(10)
	}

	for i := 0; i < 10; i++ {
		go readSum()

		go readSumSync()
	}

	//防止提前退出
	time.Sleep(time.Second * 3)
}

func add(i int) {
	sum += i
}

/*
	你期待的结果可能是“和为 1000”，但当运行程序后，可能如预期所示，
	但也可能是 990 或者 980。
	导致这种情况的核心原因是资源 sum 不是并发安全的，
	因为同时会有多个协程交叉执行 sum+=i，产生不可预料的结果
*/

func addMutex(i int) {
	mutex.Lock()
	sumSync += i
	mutex.Unlock()
}

func addMutex2(i int) {
	/*
		Mutex中的Lock\Unlock总是成对出现,而且要确保Lock获得锁后,一定执行Unlock
		所以在函数或者方法中 使用defer 语句释放锁
	*/
	mutex.Lock()
	defer mutex.Unlock()

	sumSync += i
}

/*
	对sum的加法进行了加锁,可以保证在修改sum值时,并发安全,
	但是,如果读取操作也是采用多协程的呢?
*/

func readSum() {
	fmt.Println("读写都不加锁的情况: sum is : ", sum)
	fmt.Println(" 写加锁读不加锁情况: sumSync is : ", sumSync)
}

func readSumSync() {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println(" 写不加锁,读加锁后的 情况: sum is : ", sum)
	fmt.Println(" 读,写加锁后的 情况: sumSync is : ", sumSync)

}
