package main

import "fmt"

func main() {
	var moth = 13
	fmt.Println(season(moth))
}

func season(month int) string {
	switch month {
	case 3, 4, 5:
		return "this season is spring"
	case 6, 7, 8:
		return "this season is summer"
	case 9, 10, 11:
		return "this season is autumn"
	case 12, 1, 2:
		return "this season is winter"
	default:
		return "this month is not a valid number"
	}
}
