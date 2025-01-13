package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)

	if s[0] >= 'a' && s[0] <= 'z' {
		fmt.Printf("%c", s[0]-32)
	} else {
		fmt.Printf("%c", s[0])
	}

	fmt.Printf("%s", s[1:])
}
