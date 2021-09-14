package main

import (
	"fmt" // Allows to use Println
	"os"  // Allows to use os.Stat
)

func main() {
	if _, err := os.Stat("file_exists.go"); err == nil { // _ is a blank identifier, then if error == false print exists
		fmt.Println("File exists")
	} else {
		fmt.Println("File does not exist")
	}
}
