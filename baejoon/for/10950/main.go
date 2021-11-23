package main

import "fmt"

func main() {
	var c int

	fmt.Scan(&c)
	for ; 0 < c; c-- {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
