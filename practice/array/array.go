package main

import (
	"fmt"
)

func main() {

	var anArray [5]string
	anArray[0] = "1"
	anArray[1] = "2"

	fmt.Println(anArray)

	var numbers = [6]int{0, 2, 3, 4, 5, 5}
	fmt.Println(numbers)
	fmt.Println("Size of the capacity is ", len(numbers))

}
