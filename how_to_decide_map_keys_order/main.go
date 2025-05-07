package main

import "fmt"

func main() {
	m := make(map[int]int, 10)

	m[1] = 1
	m[10] = 10
	m[8] = 8
	m[4] = 4
	m[5] = 5
	m[-1] = -1
	m[3] = 3
	m[7] = 7
	m[6] = 6
	m[9] = 9
	m[2] = 2

	fmt.Println(m)
}
