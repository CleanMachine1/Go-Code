package main

import (
	"fmt"
	"os"
	"strconv"
)

const KM_TO_MILES, MILES_TO_KM float64 = 0.62137, 1.609344

func StringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return f
}

func main() {

	var old_distance, unit string
	fmt.Println("Enter a number for distance")
	fmt.Scanln(&old_distance)
	fmt.Println("What is the unit for the number entered?\n (1)Miles\n (2)Kilometers")
	fmt.Scanln(&unit)
	var new_distance float64 = StringToFloat64(old_distance)

	switch unit {
	case "1":
		fmt.Println("Your distance in kilometers is", new_distance*MILES_TO_KM)
	case "2":
		fmt.Println("Your distance in miles is", new_distance*KM_TO_MILES)
	default:
		fmt.Println("Invalid input")
	}

}
