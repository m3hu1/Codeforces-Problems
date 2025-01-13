package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	l, r := 1, n*m

	for l < r {
		mid := (l + r) / 2
		cnt := 0
		for i := 1; i <= n; i++ {
			cnt += min(m, mid/i)
		}
		if cnt >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}

	fmt.Println(l)
}
