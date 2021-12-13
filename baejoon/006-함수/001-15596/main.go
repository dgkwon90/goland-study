//정수 N개의 합
package main

import (
	"fmt"
	"os"
	"strconv"
)

func changeNumber(a []string) []int {
	var err error
	n := make([]int, len(a))
	for i, v := range a {
		if n[i], err = strconv.Atoi(v); err != nil {
			panic(err)
		}
	}
	return n
}

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func main() {
	args := os.Args[1:]
	n := changeNumber(args)
	fmt.Println(sum(n))
}
