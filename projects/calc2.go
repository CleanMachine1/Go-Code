package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Math(num1, num2 float64, operator string) float64 {

	switch operator {
	case "1":
		return num1 + num2

	case "2":
		return num1 - num2

	case "3":
		return num1 * num2

	case "4":
		return num1 / num2

	default:
		return num1 + num2
	}
}

func StringToFloat64(str string) float64 {
	temp, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return temp
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("Simple calculator")
	fmt.Println("Supports floating points\n")

	fmt.Println("Please enter first number:")
	scan.Scan()
	stringnum1 := scan.Text()

	fmt.Println("Please enter your operator: (1)")
	fmt.Println("1. +")
	fmt.Println("2. -")
	fmt.Println("3. *")
	fmt.Println("4. /")
	scan.Scan()
	operator := scan.Text()

	fmt.Println("Please enter second number:")
	scan.Scan()
	stringnum2 := scan.Text()
	num1 := StringToFloat64(stringnum1)
	num2 := StringToFloat64(stringnum2)
	result := Math(num1, num2, operator)
	fmt.Println("Output:", result)
}
