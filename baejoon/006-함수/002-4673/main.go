//셀프 넘버
package main

import "fmt"

// func d(n int) int {
// 	r := n + (n % 10)
// 	for i := 10; i <= n; i *= 10 {
// 		r += (n / i) % 10
// 	}
// 	return r
// }

// func main() {
// 	limitNum := 10000
// 	sourceNum := make(map[int]int)
// 	dNum := make([]int, limitNum)

// 	for n := 1; n <= limitNum; n++ {
// 		sourceNum[n] = n
// 		dNum = append(dNum, d(n))
// 	}

// 	for _, v := range dNum {
// 		delete(sourceNum, v)
// 	}

// 	for i := 1; i < limitNum; i++ {
// 		if v, ok := sourceNum[i]; ok {
// 			fmt.Println(v)
// 		}
// 	}
// }

//더 좋은 방법
var notConstructorNum [10001]bool

func d(n int) int {
	r := n
	for n != 0 {
		r += n % 10
		n /= 10
	}
	return r
}
func main() {
	limitNum := 10000

	for i := 1; i <= limitNum; i++ {
		dNum := d(i)
		if dNum <= limitNum {
			notConstructorNum[dNum] = true
		}
	}

	notConstructorNum[0] = true

	for i := range notConstructorNum {
		if !notConstructorNum[i] {
			fmt.Println(i)
		}
	}
}
