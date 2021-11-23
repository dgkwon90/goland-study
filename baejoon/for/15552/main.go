//빠른 A+B

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	c := 0
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fscanln(r, &c)
	for ; c > 0; c-- {
		var a, b int
		fmt.Fscanln(r, &a, &b)
		fmt.Fprintln(w, a+b)
	}
}
