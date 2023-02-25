package main

import "fmt"

func main(){
//맵 순회 및 조회
map1 := map[string]string{
	"daum":"http://daum.net",
	"naver":"http://naver.com",
	"google":"http://google.com",
}

fmt.Println(map1["daum"])

for k,v := range map1{
	fmt.Println(k,v)
}

for _,v := range map1{
	fmt.Println(v)
}

}

