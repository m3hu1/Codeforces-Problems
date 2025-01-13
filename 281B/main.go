package main

import (
	"fmt"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)

	A, B := -1, 0
	min_diff := int64(-1)

	for b := 1; b <= n; b++ {
		a := b * x / y
		if abs(x*b-a*y) > abs((a+1)*y-x*b) {
			a++
		}
		diff := int64(abs(x*b - a*y))
		if min_diff == -1 || diff*int64(B) < min_diff*int64(b) {
			min_diff = diff
			A, B = a, b
		}

	}

	fmt.Printf("%d/%d\n", A, B)
}
