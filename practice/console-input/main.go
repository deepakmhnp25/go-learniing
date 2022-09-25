package main

/*
	This class contains sample console input based
	operations
*/
import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader((os.Stdin))
	fmt.Print("Enter text: ")
	// if you want to ignore a variable , use _
	// a method will either return a value or an exception as
	// and output.
	input, _ := reader.ReadString('\n') // denotes that when there is an enter, get the value from the console
	fmt.Println("You entered : ", input)

}
