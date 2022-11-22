package main

import (
	"fmt"
)

/** A slice uses array in the background. But unlike array slice is resizable */

func main() {
	// create an array
	var colors = [3]string{"blue", "green", "orange"}

	fmt.Println(colors)

	// if i remove the defined size, then its a slice
	var colorSlice = []string{"blue", "green", "yello", "orange"}

	// add an item to the slice
	colorSlice = append(colorSlice, "indigo")
	fmt.Println(colorSlice)

	// remove an item from the slice
	colorSlice = append(colorSlice[1:len(colorSlice)])
	fmt.Println(colorSlice)

	// if i do not mention the first index, its considered as 0
	// removing the last element
	colorSlice = append(colorSlice[:len(colorSlice)-1])
	fmt.Println(colorSlice)

	// declare a slize with type of the objects
	numbers := make([]int, 5, 5) // first 5 is the inital size, second 5 is the cap (i.e we cant add more)

	animals := make([]string, 9) // here omitting the cap makes it resizable and add more values
}
