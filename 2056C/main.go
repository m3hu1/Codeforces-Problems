package main

/*
-> 	Codeforces Template (Go)
-> 	Author: Mehul Pathak
-> 	Version: 1.2.2
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer out.Flush()

	// t := 1
	t := ri()

	// "Go"od luck!
	for t > 0 {
		n := ri()
		a := make([]int, n)

		a[1+1], a[0], a[n-1] = 1, 1, 1

		x := 0
		x++
		x++

		for i := 0; i < n; i++ {
			if a[i] == 0 {
				a[i] = x
				x++
			}
		}
		for i := 0; i < n; i++ {
			print(a[i], " ")
		}
		println()

		// test case ends here
		t--
	}
}

// cpp style typedefs
type long = int64
type float = float32
type double = float64

// faster i/o
var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

// pythonic printing
func print(arg ...interface{})            { fmt.Fprint(out, arg...) }
func println(arg ...interface{})          { fmt.Fprintln(out, arg...) }
func printf(f string, arg ...interface{}) { fmt.Fprintf(out, f, arg...) }

func scanln() string {
	x, _ := in.ReadString('\n')
	return strings.Trim(x, " \r\n")
}
func scan(s []string) string {
	if len(s) == 0 {
		return scanln()
	}
	return s[0]
}

func readLine() string          { return scanln() }                                            // readLine
func rs() string                { return scanln() }                                            // readString
func rss() []string             { return strings.Split(scanln(), " ") }                        // readStrings
func readBool(s ...string) bool { t, _ := strconv.ParseBool(scan(s)); return t }               // readBool
func readByte(s ...string) byte { t, _ := strconv.ParseUint(scan(s), 10, 8); return byte(t) }  // readByte
func rd(s ...string) double     { t, _ := strconv.ParseFloat(scan(s), 64); return t }          // readDouble
func rf(s ...string) float      { t, _ := strconv.ParseFloat(scan(s), 32); return float32(t) } // readFloat
func ri(s ...string) int        { t, _ := strconv.Atoi(scan(s)); return t }                    // readInt
func readLong(s ...string) long { t, _ := strconv.ParseInt(scan(s), 10, 64); return t }        // readLong

func readBools() []bool { // readBools
	s := rss()
	arr := make([]bool, len(s))
	for i, s := range s {
		arr[i] = readBool(s)
	}
	return arr
}
func readBytes() []byte { // readBytes
	s := rss()
	arr := make([]byte, len(s))
	for i, s := range s {
		arr[i] = readByte(s)
	}
	return arr
}
func rds() []double { // readDoubles
	s := rss()
	arr := make([]double, len(s))
	for i, s := range s {
		arr[i] = rd(s)
	}
	return arr
}
func rfs() []float { // readFloats
	s := rss()
	arr := make([]float, len(s))
	for i, s := range s {
		arr[i] = rf(s)
	}
	return arr
}
func ris() []int { // readInts
	s := rss()
	arr := make([]int, len(s))
	for i, s := range s {
		arr[i] = ri(s)
	}
	return arr
}
func readLongs() []long { // readLongs
	s := rss()
	arr := make([]long, len(s))
	for i, s := range s {
		arr[i] = readLong(s)
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
