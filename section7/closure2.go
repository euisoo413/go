//함수 Closure(2)
package main

import "fmt"

func main() {

	//예제1 함수에서 함수를 리턴한다.
	cnt := increaseCnt()

	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())

}

func increaseCnt() func() int {
	n := 0 //지역변수(캡처됨)
	return func() int {
		n += 1
		m := 1
		return m 
	}
}
