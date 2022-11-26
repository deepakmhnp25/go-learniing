package main

import (
	"fmt"
)

func main() {

	// If statements are same as java but wihout the () for conditions
	boolValue := true

	if boolValue {
		fmt.Println("The value of the booolean is true")
	}

	// in other way you can define like this
	if boolValue = true; !boolValue {
		fmt.Println("Boolean falue is false")
	}

	integerVal := 490

	if integerVal > 300 {
		fmt.Println("The value of the integer is greater than 300")
	}

}
