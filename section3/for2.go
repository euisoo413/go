package main

import "fmt"

func main(){

	//for문으로 sum 계산하기 (1~100)
	sum := 0
	for i:=0;i<=100;i++{
		sum+=i // sum = sum + i
		// sum=+1과 같은 방식으로 적으면 안됨
	}
	fmt.Println(sum)

	//외국에서는 for - sum계산을 어떤식으로 하는지
	sum2,i := 0,0
	for i<= 100 {
		sum2+=i
		i++
		//i++은 (후치연산) 반환값이 없어서 j:=i++ 같은 방식으로 계산하지 못함
	}
	fmt.Println(sum2)

	//for문으로 while문(무한루프) 구현하기
	sum3,i :=0,0
	for {
		if i>100{
			break
		}
		sum3+=i
		i++
	}
	fmt.Println(sum3)

	//i,j를 동시에 걸수 있는데 이때는 i,j의 순서를 맞춰주는 것이 좋음
	for i,j := 0,0; i<10&&j<1000; i,j= i+1,j+10{
		fmt.Println(i,j,i+j)
	}
}