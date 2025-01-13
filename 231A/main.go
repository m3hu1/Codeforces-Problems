package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)

	cnt := 0

	for i := 0; i < t; i++ {
		var p, v, t int
		fmt.Scanln(&p, &v, &t)
		if p+v+t >= 2 {
			cnt++
		}
	}

	fmt.Println(cnt)
}
