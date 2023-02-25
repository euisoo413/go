package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i is bigger than 50, smaller than 100")
	case i >= 25 && i > 50:
		fmt.Println("50>i>25")
	default:
		fmt.Println("less than 25")
	}

}
