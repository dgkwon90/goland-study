//평균
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sum float32
	var c, M int

	r := bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &c)
	for i := 0; i < c; i++ {
		score := 0
		fmt.Fscan(r, &score)
		if M < score {
			M = score
		}
		sum += float32(score)
	}
	avg := ((sum / float32(M)) * 100) / float32(c)
	fmt.Println(avg)
}
