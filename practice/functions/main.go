package main

import (
	"fmt"
)

func main() {

	multiAddition := addMultipleValues(16, 39)
	println(multiAddition)
	multiInput(12, 238, 937)
	total, length := multiReturnMethod(12, 29, 30, 30, 898)
	fmt.Println("total ", total)
	fmt.Println("length ", length)

}

// method definition with parameters and return type
func addMultipleValues(value1 int, value2 int) int {

	return value1 + value2
}

// Multiple value input using array

func multiInput(values ...int) {

	for _, v := range values {
		fmt.Println(v)
	}
}

// Method with multiple return types

func multiReturnMethod(values ...int) (int, int) {

	total := 0
	for _, v := range values {
		total += v
	}
	// return multiple values
	return total, len(values)

}
