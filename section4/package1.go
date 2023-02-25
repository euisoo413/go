//package

package main
// 

//import "fmt"
//import "os"

import (
	"fmt"
	"os"

)


func main(){

	//패키지 : 코드 구조화(모듈) 및 재사용
	//응집도, 결합도
	//Go : 패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작성할 것을 권고
	//패키지 이름 = 디렉토리 이름
	// 같은 패키지 내 소스파일들은 디렉토리명을 패키지 명으로 사용한다.
	//네이밍 규칙 : 소문자 private, 대문자 : public

	//GOROOT : 기본 패키지는 해당 폴더 위치에 있음
	//GOPATH : 사용자가 만든 패키지가 여기에 위치한다

	//Go에서 main 패키지는 특별하게 인식되어 컴파일러 공유 라이브러리가 아니다. 이는 프로그램의 시작점을 나타낸다. 
	//따라서 func main은 package main과 같은 의미가 되고 나머지 패키지들은 폴더명과 같이 쓰이게 된다.

	var name string
	fmt.Println("이름은?: ") //대문자이므로 public이다.

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi!, %s\n", name)

}