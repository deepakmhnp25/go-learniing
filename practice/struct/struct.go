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
	poodle.MethodInDogStrcut()

}

// Dog - capital letter says its public
// a type is similar to class
type Dog struct {
	Breed string
	Age   int
}

// How to create a method inside the class

func (d Dog) MethodInDogStrcut() {
	fmt.Println("This is the method inside the class Dog")
}
