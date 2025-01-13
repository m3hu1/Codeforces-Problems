package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	s1, s2 = strings.ToLower(s1), strings.ToLower(s2)

	for i := range s1 {
		if s1[i] < s2[i] {
			fmt.Println(-1)
			return
		} else if s1[i] > s2[i] {
			fmt.Println(1)
			return
		}
	}

	fmt.Println(0)
}
