package main

import "fmt"

func main() {
	const (
		a = iota // iota = 0
		b        // ioat = 1
		c        // ioat = 2
		d = "ha" // ioat = 3
		e        // ioat = 4
		f = 100  // ioat = 5
		g        // ioat = 6
		j        // ioat = 7
		h = iota // ioat = 8
		i        // ioat = 9
	)
	fmt.Println(a, b, c, d, e)
	fmt.Println(f, g, j)
	fmt.Println(h, i)
}
