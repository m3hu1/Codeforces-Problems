package main

import "fmt"

func main() {
	var l, b int
	fmt.Scan(&l, &b)

	for i := 1; ; i++ {
		l *= 3
		b *= 2
		if l > b {
			fmt.Println(i)
			break
		}
	}
}
