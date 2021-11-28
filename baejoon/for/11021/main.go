//A+B-7

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var c int

	r := bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &c)

	for i := 1; i <= c; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		fmt.Printf("Case #%v: %v\n", i, a+b)
	}
}
