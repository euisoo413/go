package main

import "fmt"

func main() {
	const (
		// _ = iota
		// A
		// B 
		// _ // 밑줄처리하면 중간에 값을 뺄수 있음
		// C

		_ = iota + 0.75 * 2
		DEFALUT
		SILVER
		_
		PLATINUM
	)
		fmt.Println(DEFALUT)
		fmt.Println(SILVER)
		fmt.Println(PLATINUM)





}
