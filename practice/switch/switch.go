package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	day := rand.Intn(7) + 1

	switch day {

	case 1:
		fmt.Println("its a Sunday")
	case 2:
		fmt.Println("its a Monday")
	case 3:
		fmt.Println("its a Tuesday")
	default:
		fmt.Println("Its some other day")
	}
}
