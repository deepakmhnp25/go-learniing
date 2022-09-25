package variables

import (
	"fmt"
)

const aConst int = 64

func VariableDeclarations() {

	var anInt int = 66
	var aString = "this is a string"
	aStringWithoutVar := "a string without var"
	var aFloat = 98.038893

	fmt.Printf("this is the type of the variable %T\n", anInt)
	fmt.Println(aString)
	fmt.Println(aConst)
	fmt.Println(aStringWithoutVar)
	fmt.Printf("the type of the aFloat variable is %T", aFloat)

}
