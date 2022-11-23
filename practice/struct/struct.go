package main

import (
	"fmt"
)

func main() {

	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf(poodle.Breed)
	// SET VALUE
	poodle.Age = 12
	fmt.Println(poodle)

}

// Dog - capital letter says its public
// a type is similar to class
type Dog struct {
	Breed string
	Age   int
}
