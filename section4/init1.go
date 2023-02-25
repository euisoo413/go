package main

import (
	"fmt"
)

func init() { //메인보다 먼저 실행됨
	// 일회성 , 프로그램 상태 체크, 변수 초기 설정 등에 쓰일 수 있음
	//init package 로드 시 가장 먼저 호출 됨
	// 가장 먼저 초기화 되는 작업 작성시 유용하다.
	fmt.Println("Init Process Start")
}

func main() {
	fmt.Println("Main Process Start")

}
