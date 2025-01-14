package main

import (
	"fmt"
	"sort"
)

// m - money, f - friendship
type F struct {
	m int
	f int
}

func count(n, d int, friends []F) int64 {
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].m < friends[j].m
	})
	var maxf, currf int64
	l := 0
	for r := 0; r < n; r++ {
		currf += int64(friends[r].f)
		for friends[r].m-friends[l].m >= d {
			currf -= int64(friends[l].f)
			l++
		}
		maxf = max(maxf, currf)
	}
	return maxf
}

func main() {
	var n, d int
	fmt.Scanln(&n, &d)

	friends := make([]F, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&friends[i].m, &friends[i].f)
	}

	result := count(n, d, friends)
	fmt.Println(result)
}
