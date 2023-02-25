//변수1

package main

import "fmt"

func main() {
	//기본 초기화
	//정수타입 : 0, 실수(소수점) : 0,0 문자열 "" boolean : true, false
	//변수명 : 숫자 첫글자 안됨, 대소문자 구분 ㅇ, 문자,숫자,밑줄, 특수기호 사용가능
	//변수 및 상수 : 함수 내외 사용 가능
	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3 //선언과 동시에 초기화
	var i float32 = 11.3
	var j string = "Hi Golang"
	var k = 4.74 //알맞는 자료형을 알아서 선택해줌
	var l = "Hi seoul"
	var m = true

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)

}
