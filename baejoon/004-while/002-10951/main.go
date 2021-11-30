// A+B - 4
package main

import (
	"fmt"
)

func main() {
	// sc := bufio.NewScanner(os.Stdin)
	// for sc.Scan() {
	// 	if s := sc.Text(); len(s) != 0 {
	// 		num := strings.Split(s, " ")
	// 		a, _ := strconv.Atoi(num[0])
	// 		b, _ := strconv.Atoi(num[1])
	// 		fmt.Println(a + b)
	// 	} else {
	// 		break
	// 	}
	// }

	// 좋은 해결 방법
	// var a, b int
	// reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)
	// defer writer.Flush()

	// for true {
	// 	val, _ := fmt.Fscanln(reader, &a, &b)
	// 	if val != 2 {
	// 		break
	// 	}
	// 	fmt.Fprintln(writer, a+b)
	// }

	// 운용
	for {
		var a, b int
		val, _ := fmt.Scanln(&a, &b)
		if val != 2 {
			break
		}
		fmt.Println(a + b)
	}
}
