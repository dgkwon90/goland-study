package main

import "fmt"

func main() {
	var m, s int

	fmt.Scan(&m)
	for i := 1; i <= m; i++ {
		s += i
	}
	fmt.Print(s)
}
