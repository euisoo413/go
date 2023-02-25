package main

import "fmt"

func main() {
	//배열
	//배열은 용량 길이 항상 같음
	//배열과 슬라이스의 차이점은 중요함
	//길이고정 vs 길이가변
	//값타입 vs 참조타입
	//복사 전달 vs 참조값 전달
	//전체 비교연산자 사용 가능 vs 비교 연산자 사용 불가
	// 대부분 슬라이스를 사용한다.

	//cap() : 배열 슬라이스 용량
	//len() : 배열 슬라이스 개수

	var arr1 [5]int // 선언안하면 그 크기만큼 다 0으로 초기화한ㄷ,
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5} //배열크기자동 맞춤
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}, //마지막엔 만드시 컴마 필요
	}

	arr1[2] = 5

	fmt.Printf("%d %v\n", len(arr1), arr1)
	fmt.Printf("%d %v\n", len(arr1), arr2)
	fmt.Printf("%d %v\n", len(arr1), arr3)
	fmt.Printf("%d %v\n", len(arr1), arr4)
	fmt.Printf("%d %v\n", len(arr1), arr5)
	fmt.Printf("%d %v\n", len(arr1), arr6)
	fmt.Printf("%d %v\n", len(arr1), arr7)

	fmt.Printf("%T\n", arr1)
	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2, len(arr1), arr2)
	fmt.Printf("%-5T %d %v\n", arr3, len(arr1), arr3)
	fmt.Printf("%-5T %d %v\n", arr4, len(arr1), arr4)
	fmt.Printf("%-5T %d %v\n", arr5, len(arr1), arr5)
	fmt.Printf("%-5T %d %v\n", arr6, len(arr1), arr6)
	fmt.Printf("%-5T %d %v\n", arr7, len(arr1), arr7)

	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10:=[...]string{"kim","lee","park"}

	fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)


}
