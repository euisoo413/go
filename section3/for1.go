// for문
// go에서 제공하는 유일한 반복문
package main

import "fmt"

func main() {
	for a := 0; a < 5; a++ {
		fmt.Println("exmaple :", a)
	}


// 예외발생1
// for a := 0; a < 5; a++ //뒤에 ;를 컴파일시 자동으로 찍기때문에 for문이 끝난걸 모름
//{
// 	fmt.Println("exmaple :", a)
// }

// 예외발생1
// for a := 0; a < 5; a++ //다른 언어처럼 대괄호 생략 안됨
// 	fmt.Println("exmaple :", a)

// 무한루프
/*
	for {
		fmt.Println("infinite loop")
	}
*/

//*** Range용법 
	loc := []string{"seoul","daegu","busan"} //key-value 형태로 저장됨
	for index,name := range loc {
		fmt.Println("ex",index," :",name)
	}

// 위 예제와는 다르게 index없이 밸류만 출력하고 싶다
	for name := range loc {
		fmt.Println("ex"," :",name) //name 적는다고 그렇게 되는게 아니고 range에서 index,value는 순서로 결정되서 첫번째자리에 아무리 name이라고 적어도 인덱스값이 나옴
	}
//iota 사용 생략하기
	for _,name := range loc {
		fmt.Println("ex"," :",name)
	}



}