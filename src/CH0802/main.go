package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	//k := 4
	fmt.Printf(" main: nums = %v,len(nums) = %d,cap(nums) = %d,&nums[0]= %v\n", nums, len(nums), cap(nums), &nums[0])
	m := make(map[string]map[string]interface{})
	m["1"] = make(map[string]interface{})
	m["1"]["2"] = 8
	fmt.Printf("m['8'] value is %v\n", m["8"])
	fmt.Printf("m['8'] == nil is  %v\n", m["8"] == nil)
	fmt.Printf("m['1'] value is %v\n", m["1"])
	fmt.Printf("m['1'] == nil is  %v\n", m["1"] == nil)

	//rotate(nums, k)
}

func rotate(nums []int, k int) {
	fmt.Printf(" input: nums = %v,len(nums) = %d,cap(nums) = %d,&nums[0]= %v\n", nums, len(nums), cap(nums), &nums[0])
	k = k % len(nums)
	t := nums[len(nums)-k:]
	fmt.Printf("t = %v,len(t) = %d,cap(t) = %d,&t[0]= %v\n", t, len(t), cap(t), &t[0])
	f := nums[0 : len(nums)-k]
	fmt.Printf("f = %v,len(f) = %d,cap(f) = %d,&f[0]= %v\n", f, len(f), cap(f), &f[0])
	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
	//nums = append(nums[0:len(nums)-k], nums[len(nums)-k:]...)
	fmt.Printf("nums = %v,len(nums) = %d,cap(nums) = %d,&nums[0]= %v\n", nums, len(nums), cap(nums), &nums[0])
	fmt.Println("rotate func nums is ", nums)
}
