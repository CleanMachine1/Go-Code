package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	input := os.Args[1:]                  // collects given arguments and 1: removes the binary name
	var1 := strings.Join(input, " ")      // convert to string rather than type string[] from args
	var2 := strings.TrimPrefix(var1, "[") // Remove square brackets from using 1: earlier
	var3 := strings.Trim(var2, "]")       // Read above ^

	cmd, err := exec.Command(`bash`, `-c`, var3).Output() // 

	if err != nil {
		log.Fatal(err)
	}

	output := string(cmd)
	fmt.Println(output)
}
