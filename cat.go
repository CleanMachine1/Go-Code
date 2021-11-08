package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file := os.Args[1:]
	output := string(file[])

	byte_text, err := ioutil.ReadFile(output)
	if err != nil {
		log.Fatal()
	}
	string_text := string(byte_text)
	fmt.Printf("%s", string_text)
}
