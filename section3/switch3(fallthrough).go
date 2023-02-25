package main

import "fmt"

func main() {

	//예제1
	switch a := 30 / 15; a {
	case 2, 4, 6:
		fmt.Println("짝수")
	case 1, 3, 5:
		fmt.Println("홀수")

	}


	//예제2 fallthrough
	// fallthrough는 switch에서 특정 조건이 만족된 다음 자동으로 break되지 않고 다음 case의 성립 여부에 관계없이 명령을 실행하게된다.
	// 예를 들어 go에 맞는 걸 출력하라고 했을 때 go 조건 케이스 뒤에 fallthrough를 입력해두면 다음 조건인 ruby 조건이 일치되지 않음에도 출력됨
	switch a:="go"; a{
	case "java":
		fmt.Println("java")
	case "go":
		fmt.Println("go")
		fallthrough
	case "ruby":
		fmt.Println("ruby")
	//fallthrough를 스위치문 마지막에 둘수가 없다
	}
}
