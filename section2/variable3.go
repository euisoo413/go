package main

import "fmt"

func main() {
	//짧은 선언
	//함수안에서만 사용(전역변수 아님), 선언 후 할당 예외 발생
	//주로 제한된 범위의 함수내에서 사용할 경우 코드 가독성을 높일 수 있음 -> 추천

	// := 를 사용해서 변경되지 않는 즉, 초기화를 두번할 수 없도록 만들 수 있음

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	//	shortVar1 := 10

	fmt.Println("shortVar1 :", shortVar1, "shortVar2", shortVar2, "shortVar3", shortVar3)

	// if문 조건문에서 아래와 같이 사용할 수 있음
	if i := 10; i < 15 {
		fmt.Println("Success")
	}

}
