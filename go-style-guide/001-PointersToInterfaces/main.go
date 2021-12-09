// 1. 타입에 대한 포인터. 이것을 "Type"으로 간주 할 수 있다?
// 2. 데이터 포인터. 저장된 데이터가 포인터일 경우, 이것은 직접적으로 저장될 수 있다.
//    만약, 저장된 데이터가 값(value)인 경우, 값에 대한 포인터가 저장된다.
package main

import "fmt"

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s S) Write1(str string) {
	s.data = str
}

func (s *S) Write2(str string) {
	s.data = str
}

func main() {
	//test1
	s := S{"test"}
	fmt.Println(s.Read())
	s.Write1("change test1")
	fmt.Println(s.Read())

	s.Write2("change test2")
	fmt.Println(s.Read())

	//test2
	sVals := map[int]S{1: {"A"}}
	fmt.Println(sVals[1].Read())
	sVals[1].Write1("test2")
	fmt.Println(sVals[1].Read())
	// 아래 코드는 컴파일 되지 않을 것
	//sVals[1].Write2("test2")

	//test3
	sPtrs := map[int]*S{1: {"B"}}
	fmt.Println(sPtrs[1].Read())
	sPtrs[1].Write1("test2")
	fmt.Println(sPtrs[1].Read())
	sPtrs[1].Write2("test3")
	fmt.Println(sPtrs[1].Read())
}
