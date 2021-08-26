package main

import (
	"fmt"     // for printing the lines
	"log"     // error related
	"os"      // for capturing argv
	"os/exec" // for running the commands
	"strings" // editing the strings such as removing square brackets and adding spaces
)

func main() {
	input := os.Args[1:]                  // collects given arguments and 1: removes the binary name
	var1 := strings.Join(input, " ")      // convert to string rather than type string[] from args
	var2 := strings.TrimSuffix(var1, "[") // Remove square brackets from using 1: earlier
	var3 := strings.TrimPrefix(var2, "]") // Read above ^

	cmd, err := exec.Command(`bash`, `-c`, var3).Output() // execute the command

	if err != nil { // error capturing
		log.Fatal(err) // report time and error status code
	}

	output := string(cmd) // Converting the byte[] to string
	fmt.Println(output)   // print what the command outputted
}
