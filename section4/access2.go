package main

import (
	"fmt"
	checkUp "section4/lib" 
	//alias는 패키지가 뭐하는지 모르기 때문에 alia로 설정해서 어떤 기능하는지 표현 가능 이번에는 checkup으로 두었음
	_ "section4/lib2" // 나중에 쓸수도 있으니 일단 두기 위해 빈 식별자 둔다.

)

func main(){
//	별칭 사용
// 빈 식별자 사용

fmt.Println("10보다 큰수?: ", checkUp.CheckNum(11))

}