package main

import "fmt"

func main() {
	var username string
	fmt.Scan(&username)

	count := 0
	set := make(map[byte]bool)

	for i := range username {
		if _, ok := set[username[i]]; !ok {
			count++
			set[username[i]] = true
		}
	}

	if count%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
