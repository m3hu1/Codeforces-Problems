package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scanln(&s)

	res := []string{}

	for i := range s {
		if i%2 == 0 {
			res = append(res, string(s[i]))
		}
	}

	sort.Strings(res)

	fmt.Print(res[0])

	for i := 1; i < len(res); i++ {
		fmt.Printf("+%s", res[i])
	}
}
