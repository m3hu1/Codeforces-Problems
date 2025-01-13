package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)

	x := 0

	for i := 0; i < t; i++ {
		var stmt string
		fmt.Scanln(&stmt)

		if stmt == "X++" || stmt == "++X" {
			x++
		} else {
			x--
		}
	}

	fmt.Println(x)
}
