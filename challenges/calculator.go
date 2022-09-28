package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
This will sum two values
exception for non numbers
user should be able to select mathematical operation
user should be able to see how long it took for the operation to complete
*/
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the operation number")
	fmt.Println("1. Add 2. Subtract 3. Multiply 4. Divide 5. Mod")
	selectedOption, _ := reader.ReadString('\n')

	// to remove delimiter
	selectedOption = selectedOption[0 : len(selectedOption)-1]

	if strings.Compare("1", selectedOption) == 0 {
		fmt.Print("Enter first number ")
		firstValue, _ := reader.ReadString('\n')
		fmt.Print("Enter first number ")
		secondValue, _ := reader.ReadString('\n')

		firstValueInt, err := strconv.ParseFloat(strings.TrimSpace(firstValue), 64)
		secondValueInt, err1 := strconv.ParseFloat(strings.TrimSpace(secondValue), 64)
		if err != nil || err1 != nil {
			fmt.Println("Enter a valid number !")
		} else {
			fmt.Println("The sum is ", firstValueInt+secondValueInt)
		}

	} else if strings.Compare("2", selectedOption) == 0 {
		fmt.Println("Subtract")
	} else if strings.Compare("3", selectedOption) == 0 {
		fmt.Println("Multiply")
	} else if strings.Compare("4", selectedOption) == 0 {
		fmt.Println("Divide")
	} else if strings.Compare("5", selectedOption) == 0 {
		fmt.Println("Mod")
	} else {
		fmt.Println("Invalid Option !")
	}

	panic("Application is exiting ...")
}
