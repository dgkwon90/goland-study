//최소, 최대
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// r := bufio.NewReader(os.Stdin)
	// w := bufio.NewWriter(os.Stdout)
	// defer w.Flush()

	// var c int
	// fmt.Fscan(r, &c)

	// var min, max int = 1000000, -1000000
	// for ; c > 0; c-- {
	// 	var n int
	// 	fmt.Fscan(r, &n)
	// 	if n < min {
	// 		min = n
	// 	}

	// 	if n > max {
	// 		max = n
	// 	}
	// }
	// fmt.Fprintf(w, "%v %v", min, max)

	// 다른 방식
	reader := bufio.NewReader(os.Stdin)
	reader.ReadBytes('\n')
	line, _ := reader.ReadBytes('\n')
	min, max := getMinMax(line)

	fmt.Printf("%d %d\n", min, max)
}

func getMinMax(line []byte) (int, int) {
	var a, op, min, max int

	min = 2000000
	max = -2000000
	op = 1

	for _, c := range line {
		fmt.Println(c)
		if c >= '0' && c <= '9' {
			a = a*10 + int(c-'0')
		} else if c == '-' {
			op = -1
		} else if c == ' ' {
			a = a * op
			if min > a {
				min = a
			}
			if max < a {
				max = a
			}

			a = 0
			op = 1
		}
	}

	a *= op
	if min > a {
		min = a
	}
	if max < a {
		max = a
	}

	return min, max
}
