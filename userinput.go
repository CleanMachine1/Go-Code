package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Type a number: ")
	fmt.Scan(&a, &b)
	fmt.Println("Your numbers are:", a, "and", b)
	/* Input needs to be seperated in either a format such as
	1 2
	or
	1
	2
	*/
}
