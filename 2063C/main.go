package main

/*
-> 	Codeforces Template (Go)
-> 	Author: Mehul Pathak
-> 	Version: 3.0
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func bs(a []int, x int) bool {
	l, r := 0, len(a)
	for l < r {
		m := l + (r-l)/2
		if a[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l < len(a) && a[l] == x
}

func build(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i + 1
	}
	return res
}

func solve() {
	// "Go"od luck!
	n := ri()

	g := make([][]int, n+1)
	d := make([]int, n+1)

	var ma, mv long

	for i := 0; i < n-1; i++ {
		u, v := ri2()
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
		d[u]++
		d[v]++
	}

	o := build(n)

	sort.Slice(o, func(i, j int) bool {
		if d[o[i]] != d[o[j]] {
			return d[o[i]] > d[o[j]]
		}
		return o[i] < o[j]
	})

	for v := 1; v <= n; v++ {
		sort.Ints(g[v])
	}

	for i := 0; i < n; i++ {
		u := o[i]
		for j := i + 1; j < n; j++ {
			v := o[j]
			if !bs(g[u], v) {
				mv = max(mv, long(d[u]+d[v]-1))
				break
			}
		}
	}

	for u := 1; u <= n; u++ {
		for _, v := range g[u] {
			if u < v {
				ma = max(ma, long(d[u]+d[v]-2))
			}
		}
	}

	println(max(mv, ma))
}

func main() {
	defer out.Flush()

	// t := 1
	t := ri()

	for t > 0 {
		solve()
		// test case ends here
		t--
	}
}

// cpp style typedefs
type long = int64
type float = float32
type double = float64

// faster i/o
var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

// pythonic printing
func print(arg ...interface{})            { fmt.Fprint(out, arg...) }
func println(arg ...interface{})          { fmt.Fprintln(out, arg...) }
func printf(f string, arg ...interface{}) { fmt.Fprintf(out, f, arg...) }

// utility functions
func rs() string { // readString
	var s string
	fmt.Fscan(in, &s)
	return s
}
func rss(n int) []string { // readStrings
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	return arr
}

func rb() bool { // readBool
	var x bool
	fmt.Fscan(in, &x)
	return x
}
func rbs(n int) []bool { // readBools
	arr := make([]bool, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	return arr
}

func ri() int { // readInt
	var x int
	fmt.Fscan(in, &x)
	return x
}
func ris(n int) []int { // readInts
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	return arr
}

func rf() float { // readFloat
	var x float
	fmt.Fscan(in, &x)
	return x
}
func rfs(n int) []float { // readFloats
	arr := make([]float, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	return arr
}

func rd() double { // readDouble
	var x double
	fmt.Fscan(in, &x)
	return x
}
func rds(n int) []double { // readDoubles
	arr := make([]double, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	return arr
}

func rl() long { // readLong
	var x long
	fmt.Fscan(in, &x)
	return x
}
func rls(n int) []long { // readLongs
	arr := make([]long, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	return arr
}

func readByte() byte { // readByte
	var x int
	fmt.Fscan(in, &x)
	return byte(x)
}
func readBytes(n int) []byte { // readBytes
	arr := make([]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	return arr
}

// extra stuff
func ri2() (int, int) { // for inputting two integers in one line
	var a, b int
	fmt.Fscan(in, &a, &b)
	return a, b
}

func ri3() (int, int, int) { // for inputting three integers in one line
	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)
	return a, b, c
}

func ri4() (int, int, int, int) { // for inputting four integers in one line
	var a, b, c, d int
	fmt.Fscan(in, &a, &b, &c, &d)
	return a, b, c, d
}

// INT_MAX and INT_MIN
const MaxUint = ^uint(0)
const intmax = int(MaxUint >> 1)
const intmin = -intmax - 1

// --------------------------------------- FUNCTIONS TO MAKE MY LIFE EASIER ---------------------------------------

// sum of all elements in an array
func _sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// lower bound - returns the index of the first element that is >= x
func _lb(a []int, x int) int {
	l, r := 0, len(a)
	for l < r {
		mid := l + (r-l)/2
		if a[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// upper bound - returns the index of the first element that is > x
func _ub(a []int, x int) int {
	l, r := 0, len(a)
	for l < r {
		mid := l + (r-l)/2
		if a[mid] <= x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
