package main

import "fmt"

func main() {

	arr1 := [5]int{1, 10, 100, 1000, 10000}



	for i := 0; i <len(arr1); i++ {

		fmt.Println("ex1 :", arr1[i])

	}
	fmt.Println()

	// 가장 많이 사용하는 예시
	arr2 := [5]int{7,77,777,7777,77777}

	for i, v := range arr2 {
		fmt.Println("ex2: ",i, v)		
	}

	for _, v := range arr2 {
		fmt.Println("ex2: ", v)		
	}

	// 이건 안됨
	for v := range arr2 {
		fmt.Println("ex2: ",v)		
	}

}