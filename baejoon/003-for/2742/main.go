//기찍 N
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inNum := 0
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &inNum)
	for i := inNum; i >= 1; i-- {
		fmt.Fprintln(w, i)
	}
}
