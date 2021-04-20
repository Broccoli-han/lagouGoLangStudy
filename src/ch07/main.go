package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能为负数")
	} else {
		return a + b, nil
	}

}
