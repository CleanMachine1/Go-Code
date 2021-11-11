package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	directory1 := os.Args[1:] // Read args from command-line, ignoring the usage of the command (./ls)
	directory2 := strings.Join(directory1, " ") // Split type string[] to string for compatability
	if directory2 == "" || directory2 == "." { // directory2 = "" or "."
		directory2 = "./" // Assign just as an easy default
	}

	morethanone := strings.Fields(directory2) // Slice directory2 into morethanone
	for _, y := range morethanone {           // y will range through morethanone and be parts of the slice
		fmt.Printf("\n%s:\n", y)
		files, err := ioutil.ReadDir(y) // Read each part of the slice
		if err != nil {                 // Error checking
			log.Fatal(err)
		}
		for _, x := range files { // Read that part of y(part of slice)'s files held
			fmt.Println(x.Name()) // https://pkg.go.dev/io/fs?utm_source=gopls#FileInfo.Name
		}
	}
}
