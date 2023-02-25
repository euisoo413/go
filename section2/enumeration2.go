package main

import "fmt"

func main() {
	const (
		A = iota
		B
		C
		
	)

	const (
		Jan = iota*10
		Feb
		Mar
		Apr
		May
		Jun	
	)

	fmt.Println(A, B, C) // iota로 시작하면 0으로 시작해서 B,C는 1개씩 증가한 값 나옴

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)
	fmt.Println(Apr)
	fmt.Println(May)
	fmt.Println(Jun)
}

