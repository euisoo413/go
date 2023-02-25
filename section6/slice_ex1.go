package main

import "fmt"

func main(){

	// append 사용하기

	s1 := []int{1,2,3,4,5}
	s2 := []int{6,7}
	s3 := []int{8,9,10,11,12}

	s2 = append(s1, s2...) // 슬라이스를 붙여서 사용할 때 ... 을 사용한다.
	s3 = append(s2, s3[0:3]...) //슬라이스 중 일부를 추출해서 붙일 때는 ["시작인덱스":"붙일 개수"] 형태로 합친다.

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	s4 := make([]int,0,5) 
	/*cap 이 5인 상태인데 이 함수를 돌리면 최대 15의 길이가 필요하다. 이때 자동으로 tmp를 만들어가면서 cap을 늘리게되는데 \
	이때 내부적으로 두배씩 증가시킨다. 부족할때 마다. 따라서 5->10->20으로 늘어난다*/
	for i:=0;i<15;i++{
		s4 = append(s4,i)
		fmt.Println(s4,len(s4),cap(s4))
	}



	// 


}