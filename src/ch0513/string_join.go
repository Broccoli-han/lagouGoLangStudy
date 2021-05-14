package main

import (
	"fmt"
	"strings"
)

func main() {
	url := strings.Join([]string{"https://translate.google.com", "GET", "test", "date", "name"}, "/")

	fmt.Println(url)
}
