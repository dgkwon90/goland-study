//N 찍기
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

	fmt.Fscanln(r, &inNum)
	for i := 1; i <= inNum; i++ {
		fmt.Fprintln(w, i)
	}
}
