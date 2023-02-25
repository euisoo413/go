package main

import "fmt"

func main() {

	arr:=[]int{1,2,1,6,4,5,2,3,4,5,5,8,5,6,1,2,1,1,4,5,}
	var count [11]int
	for i:=0;i<len(arr);i++{
		count[arr[i]]++
	}
	fmt.Println(count)

	// sort := make([]int, 0, len(arr))
	// for i:=0;i<11;i++{
	// 	for j:=0;j<count[i];j++{
	// 		sort = append(sort,i)
	// 	}
	// }

	//for문 1번만 쓰기
	//
	for i := 1 ; i < 11;i++{
		count[i] += count[i-1]
	}
	
	fmt.Println(count)
	sort := make([]int, len(arr))
	for i:=0;i<len(arr);i++{
		[arr[i]]
		
		sort = append(sort,(count[i]-count[i-1]))
	}



	fmt.Println(sort)
}

