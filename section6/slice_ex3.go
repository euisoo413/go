package main

import "fmt"


func main(){

	//copy 함수를 사용해서 슬라이스가 참조가 되는 것이 아니도록 만들 수 있다.
	//copy(복사대상, 원본) // 복사된 값이 변해도 원본은 변화없다
	//make 함수를 사용해서 복사하려는 슬라이스의 공간이 미리 확보되어야 한다 -> 예제1

	slice1 := []int{2,3,4,5,1,5,4}
	slice2 := make([]int, 5)
	slice3 := []int{}


	copy(slice2,slice1)
	copy(slice3,slice1)

	fmt.Println("ex1:",slice1)
	fmt.Println("ex1:",slice2)
	fmt.Println("ex1:",slice3)

	a := []int{1,2,3,4,5}
	b := make([]int,5)
	copy(b,a)

	b[0]=7
	b[4]=0

	fmt.Println("ex2:",a)
	fmt.Println("ex2:",b)
	
	c := [5]int{1,2,3,4,5}
	d := c[0:3] // 이렇게 바꾸면 값이 아닌 주소에 대한 참조가 됨 -> 부분적 추출은 복사가 아니다 : 진정한 복사의 슬라이싱은 아니다.

	d[0]=8
	d[2]=7

	fmt.Println("ex3:",c)
	fmt.Println("ex3:",d)

	e := []int{1,3,2,4,5,7,8,6}
	f := e[0:4:8] // 세개로 (첫값:len,끝-1:cap)을 나타낸다

	fmt.Println(len(f),cap(f),f)



}