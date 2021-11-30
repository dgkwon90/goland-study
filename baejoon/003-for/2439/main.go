//별찍기-2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// r := bufio.NewReader(os.Stdin)
	// w := bufio.NewWriter(os.Stdout)
	// defer w.Flush()

	// c := 0
	// fmt.Fscan(r, &c)
	// for i := 1; i <= c; i++ {
	// 	for k := 1; k <= c-i; k++ {
	// 		fmt.Fprint(w, " ")
	// 	}
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Fprint(w, "*")
	// 	}
	// 	fmt.Fprint(w, "\n")
	// }

	//좋은 방법
	var a int
	fmt.Scanln(&a)
	w := bufio.NewWriter(os.Stdout)
	for i := 1; i <= a; i++ {
		fmt.Fprintf(w, "%s%s\n", strings.Repeat(" ", a-i), strings.Repeat("*", i))
	}
	w.Flush()
}
