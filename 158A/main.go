package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var scores = make([]int, n+1)

	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	res := 0

	for i := 0; i < n; i++ {
		if scores[i] >= scores[k-1] {
			if scores[i] > 0 {
				res++
			}
		}
	}

	fmt.Println(res)
}
