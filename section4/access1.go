package main

import (
	"fmt"
	"section4/lib2"
)

func main() {
	//패키지 접근제어
	//변수 상수 함수 메서드 구조체 등 식별자
	//대문자 : 패키지 외부에서 접근
	//소문자 : 패키지 내부에서만 접근

	fmt.Println("100보다 큰수?: ", lib2.CheckNum1(101))
	fmt.Println("1000보다 큰수?: ", lib2.CheckNum2(101))
}
