// slice
package main

import "fmt"

func main(){
	// 배열과 비슷하지만 가변적임 -> 동적으로 크기가 늘어난다, 레퍼타입임 (참조값)
	// 성능상의 이유로 슬라이스타입을 매개변수로 전달함
	// 슬라이싱은 길이와 용량의 개념, 크기가 동적으로 할당

	// 2가지 선언 방법 1. 배열처럼, 2. make 함수 사용 make(자료형, 길이, 용량)

	// 예제1.

	var slice1 []int //슬라이스는 할당 사이즈를 지정하지 않으면 슬라이스로 인식, 적으면 배열 인식임
	slice2 := []int{} 
	slice3 := []int{1,2,3,4,5}
	slice4 := [][]int {
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	slice3[4] = 10 // 	값 수정 가능
	//slice1[0] = 1 // 얘는 불가

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1),cap(slice1),slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2),cap(slice2),slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3),cap(slice3),slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4),cap(slice4),slice4)


	var slice5 [] = make([]int, 5, 10) // make 함수를 써서 길이와 용량 설정을 할 수 있고 여기서는 길이는 5, 용량은 10으로 최대 10개 만큼 메모리에서 할당받는 것이다

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5),cap(slice5),slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6),cap(slice6),slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7),cap(slice7),slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8),cap(slice8),slice8)





}
