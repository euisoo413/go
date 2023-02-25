package main

import "fmt"

func main() {
	// 제어문 switch
	// switch 뒤 표현식 생략 가능
	// case 뒤 표현식 생략 가능
	// 자동 break때문에 fallthrouth 존재
	// type 분기 -> 값이 아닌 변수 타입으로 분기 가능

	a := 7

	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	switch b := 7; { //고랭에서는 위 예제 1보다 이런 방식을 선호함
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	switch c := "go"; c { //짧은 선언 C 해두고 c라고만 적어두면, c == 을 아래에서 확인함 , 컴파일시 반영
	case "go":
		fmt.Println("go")
	case "java":
		fmt.Println("java")
	default: // 해당되는 값 없는 전체를 뜻함
		fmt.Println("둘다 아님")
	}

	switch c := "go"; c + "lang" { //다른 값을 추가도 할 수 있음
	case "golang":
		fmt.Println("golang")
	case "java":
		fmt.Println("java")
	default: // 해당되는 값 없는 전체를 뜻함
		fmt.Println("둘다 아님")
	}

	switch i,j := 20,30 ; { //짧은 선언을 복수개로 할수 있고, ;를 사용해서 끝내놓아야함
	case i>j:
		fmt.Println("i is bigger")
	case i==j:
		fmt.Println("i equal j")
	case i<j: // 해당되는 값 없는 전체를 뜻함
		fmt.Println("j is bigger")
	}
}
