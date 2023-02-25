package main

import "fmt"


//구조체 만들기

type plane struct{
	name string
	price int
	model string
	value int
}

func price(p plane) int{
	return p.price * p.value
}

func (p plane)price2() int{
	return p.price * p.value
}

type car struct{
	name string
	price int
	value int
}

func (c car) price2() int{
	return c.price * c.value
}

func main(){
	bmw := car{name:"50",price:10000,value:2}
	benz := car{name:"230",price:20000,value:3}
	boeing := plane{price:10000,value:2}
	airbus := plane{price:20000,value:2}
	fmt.Println(bmw.price2())
	fmt.Println(benz.price2())
	fmt.Println(airbus.price2())
	fmt.Println(boeing.price2())
}