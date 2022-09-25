package main

/*
	This class contains sample console input based
	operations
*/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader((os.Stdin))
	fmt.Print("Enter text: ")
	// if you want to ignore a variable , use _
	// a method will either return a value or an exception as
	// and output.
	input, _ := reader.ReadString('\n') // denotes that when there is an enter, get the value from the console
	fmt.Println("You entered : ", input)

	// Basic string to number conversion
	fmt.Print("Enter a number: ")
	number, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(number), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value of the number : ", aFloat)
	}

}
