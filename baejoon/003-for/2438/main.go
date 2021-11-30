//별 찍기

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	c := 0
	fmt.Fscan(r, &c)
	for i := 1; i <= c; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprint(w, "*")
		}
		fmt.Fprint(w, "\n")
	}
}
