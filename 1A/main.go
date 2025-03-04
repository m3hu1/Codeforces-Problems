package main

/*
-> 	Codeforces Template (Go)
-> 	Author: Mehul Pathak
-> 	Version: 2.0.0
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve() {
	// "Go"od luck!
	n, m, a := ri3()

	println(long(math.Ceil(double(n)/double(a)) * math.Ceil(double(m)/double(a))))
}

func main() {
	defer out.Flush()

	t := 1
	// t := ri()

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
func scanln() string {
	x, _ := in.ReadString('\n')
	return strings.Trim(x, " \r\n")
}

func rs() string    { return scanln() }                     // readString
func rss() []string { return strings.Split(scanln(), " ") } // readStrings

func rb() bool { // readBool
	var x bool
	fmt.Fscan(in, &x)
	return x
}
func rbs() []bool { // readBools
	s := rss()
	arr := make([]bool, len(s))
	for i, v := range s {
		arr[i], _ = strconv.ParseBool(v)
	}
	return arr
}

func ri() int { // readInt
	var x int
	fmt.Fscan(in, &x)
	return x
}
func ris() []int { // readInts
	s := rss()
	arr := make([]int, len(s))
	for i, v := range s {
		arr[i], _ = strconv.Atoi(v)
	}
	return arr
}

func rf() float { // readFloat
	var x float
	fmt.Fscan(in, &x)
	return x
}
func rfs() []float { // readFloats
	s := rss()
	arr := make([]float, len(s))
	for i, v := range s {
		val, _ := strconv.ParseFloat(v, 32)
		arr[i] = float(val)
	}
	return arr
}

func rd() double { // readDouble
	var x double
	fmt.Fscan(in, &x)
	return x
}
func rds() []double { // readDoubles
	s := rss()
	arr := make([]double, len(s))
	for i, v := range s {
		arr[i], _ = strconv.ParseFloat(v, 64)
	}
	return arr
}

func rl() long { // readLong
	var x long
	fmt.Fscan(in, &x)
	return x
}
func rls() []long { // readLongs
	s := rss()
	arr := make([]long, len(s))
	for i, v := range s {
		arr[i], _ = strconv.ParseInt(v, 10, 64)
	}
	return arr
}

func readByte() byte { // readByte
	var x int
	fmt.Fscan(in, &x)
	return byte(x)
}
func readBytes() []byte { // readBytes
	s := rss()
	arr := make([]byte, len(s))
	for i, v := range s {
		val, _ := strconv.Atoi(v)
		arr[i] = byte(val)
	}
	return arr
}

// extra stuff
func ri2() (int, int) { // for inputting two integers in one line
	fields := strings.Fields(scanln())
	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[1])
	return a, b
}

func ri3() (int, int, int) { // for inputting three integers in one line
	fields := strings.Fields(scanln())
	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[1])
	c, _ := strconv.Atoi(fields[2])
	return a, b, c
}

func ri4() (int, int, int, int) { // for inputting four integers in one line
	fields := strings.Fields(scanln())
	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[1])
	c, _ := strconv.Atoi(fields[2])
	d, _ := strconv.Atoi(fields[3])
	return a, b, c, d
}

// INT_MAX and INT_MIN
const MaxUint = ^uint(0)
const intmax = int(MaxUint >> 1)
const intmin = -intmax - 1
