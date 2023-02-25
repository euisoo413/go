package main

import "fmt"

func main() {
	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("ex1 :", b1)
	fmt.Println("ex2 :", b2)
	fmt.Println("ex3 :", b3)

	fmt.Println("ex4 :", !b1)
	fmt.Println("ex5 :", b1 && b3)
	fmt.Println("ex6 :", b3 == b3)
	fmt.Println("ex7 :", b1 || b2)

}
