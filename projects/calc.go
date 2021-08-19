package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func divide(num1, num2 float64) float64 {
	return num1 / num2
}

func StringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return f
}

func main() {
	var choice int
	var input1, input2 string
	fmt.Printf(" 1. Add \n 2. Subtract \n 3. Multiply \n 4. Divide \n ")
	fmt.Scanln(&choice)
	fmt.Printf("Enter your 2 numbers: ") // On the same line
	fmt.Scanln(&input1, &input2)
	var float1 float64 = StringToFloat64(input1)
	var float2 float64 = StringToFloat64(input2)

	// Switch to solve the problem

	switch choice {
	case 1: // Like if choice == 1
		fmt.Println("Output =", add(float1, float2))
	case 2:
		fmt.Println("Output =", subtract(float1, float2))
	case 3:
		fmt.Println("Output =", multiply(float1, float2))
	case 4:
		fmt.Println("Output =", divide(float1, float2))
	default:
		fmt.Println("Invalid operation choice!")

	}
}
