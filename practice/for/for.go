package main

import (
	"fmt"
)

func main() {

	colors := []string{"Yellow", "Orange", "Green"}

	// definition 1
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// deinition 2
	for i := range colors {
		fmt.Println(colors[i])
	}

	//definition 3 , first position is the index, second position is the value . we ignored index by _
	for _, value := range colors {
		fmt.Println(value)
	}

	// defintion 3, with only condition

	value := 1
	for value < 10 {
		fmt.Println(value)
		value++
	}

	// break and continue will work as same as java

	// label example

	sum := 1

	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
		if sum < 200 {
			// calling the label
			goto theEndLabel
		}

	}

	// label definition
theEndLabel:
	fmt.Println("End of the program")

}
