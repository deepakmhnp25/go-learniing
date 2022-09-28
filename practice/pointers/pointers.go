package main

import (
	"fmt"
)

func main() {

	// variables which stores memory address of another variable
	anInt := 42
	var p = &anInt
	fmt.Println("Value of p : ", *p)

	*p = *p * 32
	fmt.Println("Value of p: ", *p)
	fmt.Println("Value of anInt: ", anInt)
}
