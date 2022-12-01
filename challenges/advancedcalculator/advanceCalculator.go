package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Get 2 values
	reader := bufio.NewReader((os.Stdin))

	fmt.Println("1. Add 2. Subtract 3. Multiply 4. Divide 5. Mod")
	selectedOption, _ := reader.ReadString('\n')
	selectedOption = strings.TrimSpace(selectedOption)

	fmt.Print("Enter value 1: ")
	value1, _ := reader.ReadString('\n')
	fmt.Print("Enter value 2: ")
	value2, _ := reader.ReadString('\n')

	value1Float, err1 := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	value2Float, err2 := strconv.ParseFloat(strings.TrimSpace(value2), 64)

	if err1 == nil && err2 == nil {

		operation := Operation{selectedOption, value2Float, value1Float}
		operation.SelectOperation()

	} else {
		fmt.Println("Invalid values")
	}

	// select the operation
	// display the results

	// tech implementation includes, a struct , a switch, a method of struct, a label
	// Hint : trim the white spaces of an input value

}

type Operation struct {
	Operation   string
	InputValue1 float64
	inputValue2 float64
}

func (o Operation) SelectOperation() {

	switch o.Operation {
	case "1":
		o.Addition()
	default:
		fmt.Println("Invalid operation selected")
	}

}

func (o Operation) Addition() {

	fmt.Println("The sum is : ", o.InputValue1+o.inputValue2)
}
