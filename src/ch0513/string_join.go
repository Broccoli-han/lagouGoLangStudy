package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	//url := strings.Join([]string{"https://translate.google.com", "GET", "test", "date", "name"}, "/")
	//
	//fmt.Println(url)
	dateTime := time.Now()
	yearStr, monthStr, dayStr := dateTime.Date()
	hourStr := dateTime.Hour()
	minInt := dateTime.Minute()
	minStr := fmt.Sprintf("%02d", dateTime.Minute())
	secStr := fmt.Sprintf("%02d", dateTime.Second())

	n := 2
	sInt := fmt.Sprintf("%02d", n)

	fmt.Println("date year:", yearStr)
	fmt.Println("date month:", monthStr)
	fmt.Println("date day:", dayStr)
	fmt.Println("hour:", hourStr)
	fmt.Println("minStr:", minStr)
	fmt.Println("minInt:", minInt)
	fmt.Println("second:", secStr)
	fmt.Println("sint:", sInt)
	fmt.Println("type of sint:", reflect.TypeOf(sInt))
	fmt.Println("type of n:", reflect.TypeOf(n))

}
