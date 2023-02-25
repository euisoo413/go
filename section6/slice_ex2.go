package main

import (
	"fmt"
	"sort"

)

func main(){

/*
slice[i:j] -> i부터 j-1까지 추출
slice[i:] -> i부터 끝까지
slice[:j] -> 처음부터 j-1까지
slice[:]  -> 처음부터 끝까지
*/

slice1:=[]int{3,4,1,5,6,1}
fmt.Println("ex1", slice1[1:4])
fmt.Println("ex1", slice1[1:])
fmt.Println("ex1", slice1[:4])
fmt.Println("ex1", slice1[:])


// sorting하기
slice2:=[]int{4,2,1,3,6,7,5,9,8}

fmt.Println("ex2", slice2)
fmt.Println("ex2", sort.IntsAreSorted(slice2))
sort.Ints(slice2)
fmt.Println("ex2", slice2)

slice3:=[]string{"a","c","f","g","b","d","e"}
fmt.Println("ex2", slice3)
fmt.Println("ex2", sort.StringsAreSorted(slice3))
sort.Strings(slice3)
fmt.Println("ex2", slice3)



}

