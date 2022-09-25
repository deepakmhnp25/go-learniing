package main

/*
	This class contains types of variable declarations
*/
import (
	"fmt"
)

// declaring a constant
const aConstant = 45

func main() {

	// Assign data type when initializing
	var aString string = "This is go"
	fmt.Println(aString)
	fmt.Printf("The type of the variable is %T\n", aString)

	// variable initialization without data type
	var anInt = 42
	fmt.Printf("The type of the variable is %T\n", anInt)

	// using := to declare variables
	// := works only inside a function
	anotherString := "This is another string"
	fmt.Printf("The type of the variable is %T\n", anotherString)

	fmt.Printf("The type of the constant is %T\n", aConstant)
	fmt.Println(aConstant)
}
