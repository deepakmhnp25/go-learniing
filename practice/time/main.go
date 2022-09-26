package main

import (
	"fmt"
	"time"
)

func main() {
	// print current time
	currentTime := time.Now()
	fmt.Println("The current time is ", currentTime)

	// print formatted time
	fmt.Println("The formatted time is", currentTime.Format(time.ANSIC))

	// create a custom time object

	customTime := time.Date(2022, time.September, 26, 20, 52, 37, 0, time.UTC)
	fmt.Println("The custom time is ", customTime)

	parsedTime, _ := time.Parse(time.ANSIC, "Mon Sep 26 19:54:02 2022")

	fmt.Println("The parsed time is ", parsedTime)
}
