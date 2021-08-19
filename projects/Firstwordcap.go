package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter a word")
	fmt.Scanln(&input)
	fmt.Println(strings.Title(input)) //strings also has ToTitle which caps the entire string.

}
