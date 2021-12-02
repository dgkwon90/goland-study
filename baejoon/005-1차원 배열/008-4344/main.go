//평균은 넘겠지

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	testCaseCount := 0
	N := 0

	fmt.Fscan(reader, &testCaseCount)

	for i := 0; i < testCaseCount; i++ {
		fmt.Fscan(reader, &N)
		scores := make([]int, N, N)
		var sum float32
		for j := 0; j < N; j++ {
			score := 0
			fmt.Fscan(reader, &score)
			sum += float32(score)
			scores = append(scores, score)
		}

		avg := sum / float32(N)
		goodScore := 0
		for _, score := range scores {
			if avg < float32(score) {
				goodScore++
			}
		}
		fmt.Printf("%.3f%%\n", float32(goodScore)/float32(N)*100)
	}
}
