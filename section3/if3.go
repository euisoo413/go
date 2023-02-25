// if문3
// if - else if -else

package main

import "fmt"

func main(){
	i:=100

	if i<=30{
		fmt.Println("a는 30 이하")
	} else if i>30 && i<= 50{
		fmt.Println("a는 30 이상 50 이하")
	} else {		
		fmt.Println("a는 50 이상")
	}
}
