package main

import "fmt"

func main() {
	var m, s int

	fmt.Scan(&m)
	for i := 1; i <= m; i++ {
		s += i
	}
	fmt.Print(s)
	// 더 좋은 방법
	// to := 0
	// fmt.Scan(&to)
	// fmt.Println(to * (to + 1) / 2)
}
