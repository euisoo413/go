package main

import "fmt"

type Stringer interface{
	String() string	
}

func (s Student) String() string{
	return fmt.Sprintf("%d",s.Age)
}

type Student struct{
	Age int
	Name string
}

func main(){
	student1 := Student{10,"sam"}
	var stringer Stringer
 	stringer = student1
	 fmt.Println(student1.String()) 
	fmt.Println(stringer.String())

}

