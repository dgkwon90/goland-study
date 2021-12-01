//숫자의 개수
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	sum := strconv.Itoa(a * b * c)

	for i := 0; i < 10; i++ {
		fmt.Println(strings.Count(sum, strconv.Itoa(i)))
	}
}
