package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanln(&n, &k)

	t, res := 0, 0
	l, r := 0, n

	for l <= r {
		mid := l + (r-l)/2
		t = 5 * mid * (mid + 1) / 2
		if t > 240-k {
			r = mid - 1
		} else {
			l = mid + 1
			res = mid
		}
	}

	fmt.Println(res)
}
