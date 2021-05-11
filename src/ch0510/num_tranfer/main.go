package main

import (
	"fmt"
	"math"
)

type ByteSize float64

const (
	_           = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	//a := IntFromFloat64(12.0123)
	//print(a)
	fmt.Printf("KB value is %v\n", KB)
	fmt.Printf("MB value is %v\n", MB)
	fmt.Printf("GB value is %v\n", GB)
	fmt.Printf("TB value is %v\n", TB)
	fmt.Printf("PB value is %v\n", PB)
	fmt.Printf("EB value is %v\n", EB)
	fmt.Printf("ZB value is %v\n", ZB)
	fmt.Printf("YB value is %v\n", YB)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
		whole, fraction := math.Modf(x)
		print("whole is : ", whole)
		print("fraction is : ", fraction)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
