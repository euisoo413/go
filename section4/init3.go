package main

import (
	"fmt"
	"section4/lib"
)


var num int32
func init(){

	num=30
	fmt.Println("main inti start")
}

func main(){
	fmt.Println("10보다 큰수?:", lib.CheckNum(num))
}
