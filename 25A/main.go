package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	nums := make([]int, n)
	even, odd := 0, 0
	even_idx, odd_idx := 0, 0

	for i := range nums {
		fmt.Scan(&nums[i])
		if nums[i]%2 == 0 {
			even++
			even_idx = i
		} else {
			odd++
			odd_idx = i
		}
	}

	if even > odd {
		fmt.Println(odd_idx + 1)
	} else {
		fmt.Println(even_idx + 1)
	}
}
