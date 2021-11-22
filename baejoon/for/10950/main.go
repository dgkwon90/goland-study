package main

import "fmt"

func main() {
	var c, a, b int
	var result []int

	fmt.Scan(&c)
	for i := 0; i < c; i++ {
		fmt.Scan(&a, &b)
		result = append(result, a+b)
	}

	for _, value := range result {
		fmt.Println(value)
	}
}
