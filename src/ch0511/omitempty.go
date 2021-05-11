package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func main() {
	p1 := Person{
		Name: "zhangsan",
		Age:  19,
	}

	data, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
	fmt.Println(p1.Name, p1.Age, p1.Addr)
	p2 := Person{
		Name: "Lisi",
		Age:  20,
		Addr: "北京",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data2)
	fmt.Println(p2.Name, p2.Age, p2.Addr)
}
