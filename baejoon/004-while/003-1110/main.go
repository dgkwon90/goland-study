// 더하기
// 규칙을 찾자.
// 55
// 1) 5+5 10
// 50
// 2) 5+0 5
// 05
// 3) 0+5 5
// 55

package main

import (
	"fmt"
)

func main() {
	// r := bufio.NewReader(os.Stdin)
	// w := bufio.NewWriter(os.Stdout)
	// defer w.Flush()

	// var sourceNum, newNum, cicle int
	// fmt.Fscan(r, &sourceNum)
	// newNum = sourceNum
	// for {
	// 	cicle++
	// 	a := newNum / 10
	// 	b := newNum % 10
	// 	c := a + b
	// 	newNum = ((newNum % 10) * 10) + (c % 10)
	// 	if newNum == sourceNum {
	// 		break
	// 	}
	// }
	// fmt.Fprint(w, cicle)

	//좋은 소스
	var N int
	fmt.Scan(&N)

	var t, cnt int = -1, 0
	for N != t {
		if cnt == 0 {
			t = N
		}
		t = (t%10)*10 + ((t/10)+(t%10))%10
		cnt++
	}
	fmt.Println(cnt)
}
