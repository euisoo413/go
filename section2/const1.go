package main

import "fmt"

func main() {

	//const 사용 시 바로 초기화 해야함, 한번 선언 후 값 변경 금지, 고정된 값 관리용
	//
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	//const d = getHeight()
	//항상 똑같은 값을 리턴한다는 보장이 없어서 에러남 
	const e = 35.6
	const f = false

	/*
	에러발생 
	const g string
	g = "test3"
	*/

	fmt.Println("a : ", a, "b :",b, "c: ", c, "e :",e,"f :",f)

}
