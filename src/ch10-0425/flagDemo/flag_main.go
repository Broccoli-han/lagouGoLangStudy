package main

import (
	"flag"
	"fmt"
	"sort"
)

var (
	intflag    int
	boolflag   bool
	stringflag string
)

func init() {
	flag.IntVar(&intflag, "intflag", 0, "int flag value")
	flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	flag.StringVar(&stringflag, "stringflag", "default", "string flag value")
}

func main() {
	//flag.Parse()
	//
	//fmt.Println("int flag:", intflag)
	//fmt.Println("bool flag:", boolflag)
	//fmt.Println("string flag:", stringflag)

	var array = []int{4, 2, 5, 1, 0, 6}
	var arr []int
	arr = append(arr, 1, 2)
	fmt.Println(arr)
	fmt.Println(array)
	sort.Ints(array)
	fmt.Println("排序后的数组为:", array)
	fmt.Println("排序后的数组第一个元素为:", array[0])
	fmt.Println("排序后的数组去掉第一个元素为:", array[1:])
}
