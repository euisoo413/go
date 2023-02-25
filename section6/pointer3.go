package main

import "fmt"

func rptc(n *int){
	*n = 777
}

func vptc(n int){
	n = 777
}

func main(){
	//*** 함수에서는 매서드 호출 시 값을 복사해서 전달하는데, 함수 메서드 내에서는 원본 값에 대한 변경은 불가능하다.
	//원본값 변경을 위해서는 포인터로 전달을 해야함
	//큰 배열 값은 그래서 복사해서 옮겨야하는 시간 손실로 포인터로 참조해서 옮기는데, 고는 그냥 맵이나 슬라이스로 하면 편함
	a := 10
	b := 10
	
	rptc(&a) //a의 주소값으로 전달해야 오류 안난다.
	vptc(b)


fmt.Println(a)
fmt.Println(b)

}