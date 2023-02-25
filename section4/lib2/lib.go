package lib2

func CheckNum1(c int32) bool { //예를 들면 CheckNum 을 checkNum으로 바꾸면 퍼블릭에서 프라이빗으로 바껴서 딴 패키지에서 사용못함
	return c > 100
}

func CheckNum2(c int32) bool { //예를 들면 CheckNum 을 checkNum으로 바꾸면 퍼블릭에서 프라이빗으로 바껴서 딴 패키지에서 사용못함
	return c > 1000
}