package main

import "fmt"

func main(){
	arr1 := [3]int{1,2,3}
	var arr2 [3]int

	arr2 = arr1
	arr2[1] = 7

	fmt.Println("ex1:", arr1)
	fmt.Println("ex2:", arr2)


	slice1 := []int{1,2,3}
	var slice2 []int
	
	slice2 = slice1
	slice2[0] = 7

	fmt.Println("ex1:", slice1)
	fmt.Println("ex2:", slice2)


	slice3 := make([]int, 50,100)
	fmt.Println("ex3:",slice3[0])
	fmt.Println("ex3:",slice3[50])

	// 길이 인덱스 오버





}


