package main

import "fmt"

func main() {
	var wt int
	fmt.Scanln(&wt)

	if wt%2 == 0 && wt > 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
