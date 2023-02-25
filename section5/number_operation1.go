package main

import (
	"fmt"
	"math"
)
func main(){

//타입이 같아야 계산을 할 수 있음 . 특히 돈 관련 casting 시 유실이 발생할 수 있으므로
//다른타입끼리는 반드시 형 변환 필요 / 비교 연산자도 같아아함
//int 기본이며 int8,16,32등으로 사이즈 조절 가능
//형 변환 없을 경우 예외 발생
// ^: 1의 보수

var n1 unit8 = math.MaxUint8
var n2 unit16 = math.MaxUint16
var n3 unit32 = math.MaxUint32
var n4 unit64 = math.MaxUint64
fmt.Println("n1 :", n1)


n5 := 100000
n6 := int16()
n7 := uint8(300)//error 발생


}