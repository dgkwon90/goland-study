// 구구단
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for i := 1; i <= 9; i++ {
		fmt.Println(a, "*", i, "=", a*i)
	}
}
