//라이브러리 접근제어(1)

package lib
import "fmt"
func init() {

	fmt.Println("lib package > init start!")
}

func CheckNum(c int32) bool { //예를 들면 CheckNum 을 checkNum으로 바꾸면 퍼블릭에서 프라이빗으로 바껴서 딴 패키지에서 사용못함
	return c > 10
}
