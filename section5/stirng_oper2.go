package main

import "fmt"

func main(){
//문자열의 추출 조합 비교
//파이썬 처럼 자르고 붙이기가 편함


str1 := "Golang"
str2 := "world"

fmt.Println("ex1",str1 == str2)
fmt.Println("ex1",str1 != str2)
fmt.Println("ex1",str1 > str2) //결과값이 str1이 큰게 트루 같지만 반대의 결과가 나오는데 그 이유는 값 하나씩 아스키코드 비교를 하므로
fmt.Println("ex1",str1 < str2)



}