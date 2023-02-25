package main

import "fmt"

func main() {

	fmt.Println("ex1 :", true && true)
	fmt.Println("ex1 :", true && false)
	fmt.Println("ex1 :", false && false)
	fmt.Println("ex1 :", true || true)
	fmt.Println("ex1 :", true || false)
	fmt.Println("ex1 :", false || false)
	fmt.Println("ex1 :", !true)
	fmt.Println("ex1 :", !false)

	num1 := 39
	num2 := 28

	fmt.Println("ex1 :", num1 > num2)
	fmt.Println("ex1 :", num1 < num2)
	fmt.Println("ex1 :", num1 >= num2)
	fmt.Println("ex1 :", num1 <= num2)
	fmt.Println("ex1 :", num1 == num2)
	fmt.Println("ex1 :", num1 != num2)

}
