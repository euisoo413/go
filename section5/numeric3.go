package main

import "fmt"

func main() {
	//실수(부동소수점)
	//float32(7자리) , float64(15자리)

	var num1 float32 = 0.14

	var num4 float32 = 10.0

	fmt.Println("num4:",num4-0.1)
	fmt.Println("num4:",float32(num4-0.1))
	fmt.Println("num4:",float64(num40.1))
	//부동소수점 오차 해결 자릿수가 늘어남으로서 오차가 발생함
