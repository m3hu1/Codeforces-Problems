package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)

	for n := range t {
		fmt.Scanln(&n)
		cows := n / 4
		rem := n % 4
		chickens := rem / 2
		res := cows + chickens
		fmt.Println(res)
	}

	return
}
