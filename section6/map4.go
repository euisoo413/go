package main

import "fmt"

func main() {
	// 맵 조회 시 주의점
	map5 := map[string]int{
		"apple":  41,
		"orange": 26,
		"banana": 0,
	}

	value1 := map5["apple"]
	value2 := map5["banana"]
	value4 := map5["kiwi"]
	value3, ok3 := map5["banana"]
	value5, ok5 := map5["kiwi"] //키위는 없는데 인티저의 없을 때의 값이 0 이기 때문에 구분이 필요함 int:0 string: "" float : 0.0

	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value4)
	fmt.Println(value3, ok3) //바나나는 실제로 값이 0 이므로 존재 유무를 확인해야 함
	fmt.Println(value5, ok5)

	if value, ok := map5["kiwi"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Not exist")
	}

	if _, ok := map5["kiwi"]; !ok {
		fmt.Println("not exist")
	}

}
