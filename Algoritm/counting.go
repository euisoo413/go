package main

import "fmt"

func main(){

	arr:=[]int{1,3,1,5,6,7,8,4,3,2,4,1,9,8,10,2,2,4,5,6}

	var count []int
	for i:=0;i<len(arr);i++{
		count[arr[i]]++

	}



}