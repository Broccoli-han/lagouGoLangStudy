package entity

type Person struct {
	Name string
	Age  uint
}

/*结构体中的字段可以是任意类型，包括自定义的结构体类型*/
type Person2 struct {
	Name string
	Age  uint
	Addr Address
}

type Address struct {
	Province string
	City     string
}
