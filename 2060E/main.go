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
)

type Edge struct {
	u, v int
}

type Graph struct {
	n     int
	adj1  [][]int
	adj2  [][]int
	mark  []bool
	mark2 []bool
	col   []int
	mp    map[Edge]int
	cnt   int
	cnt2  int
}

func ng(n int) *Graph {
	return &Graph{
		n:     n,
		adj1:  make([][]int, n),
		adj2:  make([][]int, n),
		mark:  make([]bool, n),
		mark2: make([]bool, n),
		col:   make([]int, n),
		mp:    make(map[Edge]int),
		cnt:   0,
		cnt2:  0,
	}
}

func (g *Graph) e1(u, v int) {
	v -= 1
	u -= 1
	g.adj1[u] = append(g.adj1[u], v)
	g.adj1[v] = append(g.adj1[v], u)
}

func (g *Graph) e2(u, v int) {
	v--
	u--
	g.adj2[u] = append(g.adj2[u], v)
	g.adj2[v] = append(g.adj2[v], u)
}

func (g *Graph) helper(u int) {
	g.col[u] = g.cnt
	g.mark2[u] = true
	for _, v := range g.adj2[u] {
		if !g.mark2[v] {
			g.helper(v)
		}
	}
}

func (g *Graph) sos(u int) {
	g.mark[u] = true
	for _, v := range g.adj1[u] {
		if !g.mark[v] {
			if _, exists := g.mp[Edge{u, v}]; !exists || g.mp[Edge{u, v}] != 1 {
				g.sos(v)
			}
		}
	}
}

func (g *Graph) psg() {
	for i := 0; i < g.n; i++ {
		if !g.mark2[i] {
			g.cnt++
			g.helper(i)
		}
	}
}

func (g *Graph) cebc() int {
	ans := 0
	for i := 0; i < g.n; i++ {
		for _, v := range g.adj1[i] {
			if g.col[i] != g.col[v] {
				ans++
				g.mp[Edge{i, v}] = 1
			}
		}
	}
	return ans
}

func (g *Graph) pfg() {
	for i := 0; i < g.n; i++ {
		if !g.mark[i] {
			g.cnt2++
			g.sos(i)
		}
	}
}

func (g *Graph) Solve() int {
	g.psg()
	ans := g.cebc()
	g.pfg()
	return ans/2 + (g.cnt2 - g.cnt)
}

func solve() {
	// "Go"od luck!
	n, m1, m2 := ri3()
	graph := ng(n)

	for i := 0; i < m1; i++ {
		u, v := ri2()
		graph.e1(u, v)
	}

	for i := 0; i < m2; i++ {
		u, v := ri2()
		graph.e2(u, v)
	}

	println(graph.Solve())
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
