// 배열 : 값 복사 해보기



package main

import "fmt"

func main(){	
arr1:=[5]int{1,10,100,1000,10000}
arr2:=arr1

//확인해봐야 할 사항 : arr2가 arr1의 값을 보여주는 지 아니면 새롭게 값을 저장해서 본연의 값을 갖게되는지 확인한다.

fmt.Printf("ex1: %p %v\n", &arr1, arr1 )
fmt.Printf("ex2: %p %v\n", &arr2, arr2 )

}