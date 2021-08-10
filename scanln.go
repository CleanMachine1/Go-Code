package main

import (
	"fmt"
)

func main() {

	var a, b string

	fmt.Print("Type 2 strings: ")
	fmt.Scanln(&a, &b)
	fmt.Println("Your strings are: ", a, "and", b)
	/* Format needs to be
	String1 String2
	*/
}
