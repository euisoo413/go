// if문
// 제어문(조건문)
// 반드시 Boolean 검사를 수행하므로 True, False 만 구분하며 (0,1) 구분하지 않음
// 소괄호 사용하지 않음


package main

import "fmt"

func main(){

	var a int  = 20
	b := 20

	if a>15 {

		fmt.Println("a는 15 이상 ")
	}

	if b > 25 {

		fmt.Println("b is bigger than 25")
	}
	
	//에러발생1
	// if a>15 
	// {

	// 	fmt.Println("a는 15 이상 ")
	// }

	//에러발생2
	// if a>15 
	// 	fmt.Println("a는 15 이상 ")

	//에러발생2
	// if c := 1; c {
	// 	fmt.Println("c는 True임")
	// }
	/* 
	₩₩₩ 
	결과값 : # command-line-arguments
	 ./if1.go:38:11: non-boolean condition in if statement
	 ₩₩₩
	 */

	 //에러발생 4
	// if c := 45; c {
	// 	fmt.Println("c는 True임")
	// }
	// fmt.Println(c) -> 에러발생함 c는 if문 안에서만 사용

}