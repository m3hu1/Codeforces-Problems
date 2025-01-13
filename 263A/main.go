package main

import "fmt"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	mat := [5][5]int{}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&mat[i][j])
			if mat[i][j] == 1 {
				fmt.Println(abs(i-2) + abs(j-2))
			}
		}
	}
}
