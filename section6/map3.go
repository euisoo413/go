package main

import "fmt"

func main(){
//맵 수정 및 삭제

map1 := map[string]string{
	"daum":"http://daum.net",
	"naver":"http://naver.com",
	"google":"http://google.com",
	"home":"http://test1.com",
}

fmt.Println(map1)

map1["home2"]="http://test2.com"
fmt.Println(map1)

map1["home2"]="http://test2-2.com"
fmt.Println(map1)

//삭제 delete(map이름, )
delete(map1,"home2")
fmt.Println(map1)



// for k,v := range map1{
// 	fmt.Println(k,v)
// }

// for _,v := range map1{
// 	fmt.Println(v)
// }

}

