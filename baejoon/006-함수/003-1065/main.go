//한수
package main

import "fmt"

func arrayNum(num int) []int {
	array := make([]int, 4)
	for num != 0 {
		array = append(array, num%10)
		num /= 10
	}
	return array
}

func hanNum(inputNumArray []int) int {
	var count int
	return count
}

func main() {
	var inputNum int
	fmt.Scan(&inputNum)

	if inputNum < 10 {
		fmt.Println(1)
	} else {
		fmt.Println(hanNum(arrayNum(inputNum)))
	}
}
