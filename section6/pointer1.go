package main

import "fmt"

func main(){
	//포인터
	//변수의 지역성을 표기하며, 메모리 값이 연속될 때의 참조를 위해, 힙, 스택 드

	//파이썬 자바 등은 포인터를 지원하지 않지만 씨언어 고 언어에서는 일부 포인터를 지원한다.
	//고는 리스크가 있는 연산은 하지 못하도록 막아놨다.
	//주소 값은 변경 불가하다
	//에스터리스크 사용

	// nil로 초기화 된다.
	var a *int //이런 식으로 초기화하면 nil로 초기화됨
	var b *int =new(int) //이렇게하면 주소값을 가지고 있다

fmt.Println(a)
fmt.Println(b)

i :=7
a = &i
b = &i

fmt.Println(i, a, b, *a, *b)
i++
fmt.Println(i, a, b, *a, *b)
*a++
fmt.Println(i, a, b, *a, *b)
*b=777
fmt.Println(i, a, b, *a, *b)

}