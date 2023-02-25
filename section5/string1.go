package main

import (
	"fmt"
	"unicode/utf8"
	// "unicode/utf8"
)

func main(){
// 문자열

// 큰 따옴표 "" , 백스쿼트 
// 고랭은 문자열 char 타입이 존재하지 않음 rune(int32)으로 문자 코드 값 표현

//문자 : ' '로 작성
//자주 사용하는 escape : \\, \`, \" , \a(콘솔벨) \b(백스페이스), \f (쪽바꿈) \n 줄바꿈 \r(복귀 캐리지리턴) \t`

// 예제1
//var str1 string = "c:\go_study\src" //제대로 인식하지 못해서 아래처럼 변경
var str1 string = "c:\\go_study\\src\\" //제대로 인식하지 못해서 아래처럼 변경
str2 := `c:\go_study\src\` // 백스쿼트를 사용하면 문자를 그대로 인식한다.

fmt.Println("ex1: ", str1)
fmt.Println("ex1: ", str2)

// str3 ~ 5 한글 출력하기 + 유니코드

var str3 string = "Hello World!"
var str4 string = "안녕하세요"

//예제 바이트 길이 수
fmt.Println("ex1: ", len(str3))
fmt.Println("ex1: ", len(str4))

//실제 길이는 rune이나 유니코드 패키지를 활용해서 세야함
fmt.Println("ex1: ", utf8.RuneCountInString(str3))
fmt.Println("ex1: ", utf8.RuneCountInString(str4))
fmt.Println("ex1: ", []rune(str3))
fmt.Println("ex1: ", []rune(str4))
fmt.Println("ex1: ", len([]rune(str3)))
fmt.Println("ex1: ", len([]rune(str4)))


}