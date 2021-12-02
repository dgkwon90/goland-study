//OX
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
	for i := 0; i < c; i++ {
		score := 0
		if line, _, _ := r.ReadLine(); len(line) > 0 {
			str := string(line)
			x := 1
			for j := 0; j < len(line); j++ {

				if "O" == str[j:j+1] {
					score += x
					x++
				} else {
					x = 1
				}
			}
			fmt.Println(score)
		} else {
			i--
		}
	}
}
