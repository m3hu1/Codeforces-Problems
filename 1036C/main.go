package main

import "fmt"

func helper(n int) int {
	digits := make([]int, 0)
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n = n / 10
	}
	m := len(digits)
	cache := make(map[[3]int]int)
	var dp func(pos int, t bool, nz int) int
	dp = func(pos int, t bool, nz int) int {
		if nz > 3 {
			return 0
		}
		if pos == m {
			return 1
		}
		key := [3]int{pos, 0, nz}
		if t {
			key[1] = 1
		}
		if val, ok := cache[key]; ok {
			return val
		}
		l := 9
		if t {
			l = digits[pos]
		}
		res := 0
		for d := 0; d <= l; d++ {
			nt := t && (d == l)
			ncnt := nz
			if d != 0 {
				ncnt++
			}
			res += dp(pos+1, nt, ncnt)
		}
		cache[key] = res
		return res
	}
	return dp(0, true, 0)
}

func main() {
	var t int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		var L, R int
		fmt.Scanln(&L, &R)
		fmt.Println(helper(R) - helper(L-1))
	}
}
