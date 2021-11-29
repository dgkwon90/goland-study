//X보다 작은 수
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nc, x int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &nc, &x)
	for i := 0; i < nc; i++ {
		var in int
		fmt.Fscan(r, &in)
		if in < x {
			fmt.Fprintf(w, "%d ", in)
		}
	}
}
