package main

import (
	"fmt"
	"math"
)

func main() {

	var aInt int = 24
	aFloat := 24.73

	// rounding a float value
	roundedValue := math.Round(aFloat)
	fmt.Println(aInt)
	fmt.Printf("roundedValue data type is %T\n", roundedValue)

	// rounding a float value to nearest 2 digits

	// here == is used as the roundedValue is already defined
	roundedValue = math.Round(roundedValue*100) / 100
	fmt.Println("roundedValue data now", roundedValue)

	// circle diameter
	circleRadius := 78.99

	circleDiameter := circleRadius * 2 * math.Pi
	fmt.Println("The diameter of the circle is ", circleDiameter)

}
