//나머지
package main

import "fmt"

func main() {
	var diff map[int]int
	diff = make(map[int]int)

	for i := 0; i < 10; i++ {
		var num int
		fmt.Scan(&num)
		diff[num%42]++
	}
	fmt.Println(len(diff))
}
