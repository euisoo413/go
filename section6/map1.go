package main

import "fmt"

func main(){
	// map : 해시테이블, 딕셔너리(파이션), 키-밸류 형태로 저장
	// 레퍼런스 타입 (참조값 전달)
	// 비교 연산자 사용 불가능 (참조 타입 이므로)
	// 특징 : 참조 타입을 키로 사용 불가능하며, 모든 타입을 밸류로 사용하는건 가능하다.
	// make 함수 및 리터럴(축약) 형태로 초기화 가능
	// 순서가 없어서 뭐가 먼저 나올지 모름 (슬라이스, 배열 등과는 다름)

	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int) //자료형 생략
	map3 := make(map[string]int) // 리터럴

	fmt.Println("ex1: ",map1)
	fmt.Println("ex1: ",map2)
	fmt.Println("ex1: ",map3)

	map4 := map[string]int{}//json6 형태로 전달됨 {} 형태로 적어두면 데이터 타입에 대해 정의해줄수 있음
	map4["apple"]=40
	map4["orange"]=25
	map4["banana"]=20

	map5 := map[string]int{
		"apple":41,
		"orange":26,
		"banana":21,
	}
	map6 := make(map[string]int,5)
	map6["apple"]=40
	map6["orange"]=25
	map6["banana"]=20

	fmt.Println("ex1: ",map4)
	fmt.Println("ex1: ",map5)
	fmt.Println("ex1: ",map6)
	fmt.Println("ex1: ",map6["orange"])
	fmt.Println("ex1: ",map6["banana"])
	fmt.Println("ex1: ",map6["apple"])

}