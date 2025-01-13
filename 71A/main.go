package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		var word string
		fmt.Scanln(&word)
		if len(word) > 10 {
			fmt.Printf("%c%d%c", word[0], len(word)-2, word[len(word)-1])
			fmt.Println()
		} else {
			fmt.Println(word)
		}
	}
}
