package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var s string
	fmt.Scanln(&s)

	ans := 0

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			ans++
		}
	}

	fmt.Println(ans)
}
