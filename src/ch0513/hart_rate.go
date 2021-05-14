package main

import (
	"fmt"
	"time"
)

var (
	age       int
	heartRate int
)

func main() {
	//inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请以此输入您的年龄和静息心率（以空格符号 分隔），来计算您的最佳有氧心率区间！")
	//input, err := inputReader.Read('\n')
	//fmt.Scanln("%d, %d", &age, &heartRate)
	fmt.Scanln(&age, &heartRate)
	minRate, maxRate := optimalAerobicHeartRate(age, heartRate)
	fmt.Printf("您的最佳有氧心率区间为：[%0.1f ~ %0.1f]\n", minRate, maxRate)
	fmt.Println("程序将要在5秒后自动退出！")
	time.Sleep(5 * time.Second)
}

func optimalAerobicHeartRate(age, heartRate int) (minRate, maxRate float64) {
	minRate = float64((220-age-heartRate))*0.4 + float64(heartRate)
	maxRate = float64((220-age-heartRate))*0.6 + float64(heartRate)
	return
}
