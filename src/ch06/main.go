package main

import (
	"fmt"
	"lagouGoLangStudy/src/ch06/entity"
)

func main() {
	// 第一种声明结构，并赋值
	var p entity.Person
	p.Name = "broccoli"
	p.Age = 30

	// 第二种 结构体字面量的方式初始化
	p1 := entity.Person{"zhangsan", 20}

	//第三种 不按照结构体的顺序初始化
	p2 := entity.Person{Age: 18, Name: "Lisi"}

	// 还可以只初始化一个结构体中的一个字段, 此时其他字段是默认值的
	p3 := entity.Person{Name: "wangwu"}

	// 结构体中内嵌结构的初始化
	p4 := entity.Person2{
		Age:  19,
		Name: "zhaoliu",
		Addr: entity.Address{
			Province: "shanghai",
			City:     "shanghai",
		},
	}

	fmt.Printf("p name: %v;age: %d\n", p.Name, p.Age)
	fmt.Printf("p1 name: %v;age: %d\n", p1.Name, p1.Age)
	fmt.Printf("p1 name: %v;age: %d\n", p2.Name, p2.Age)
	fmt.Printf("p1 name: %v;age: %d\n", p3.Name, p3.Age)
	fmt.Printf("p4 name: %v;age: %d; address: %v - %v\n", p4.Name, p4.Age, p4.Addr.Province, p4.Addr.City)

}
