//최댓값

package main

import (
	"fmt"
)

func main() {
	//	r := bufio.NewReader(os.Stdin)

	var max, maxIndex = -1, 0
	//nums := make([]int, 9, 9)
	for i := 1; i <= 9; i++ {
		var num int
		fmt.Scan(&num)
		//	nums = append(nums, num)
		if num > max {
			max = num
			maxIndex = i
		}
	}

	fmt.Printf("%v\n%v", max, maxIndex)
}
