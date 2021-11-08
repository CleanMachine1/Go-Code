package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	files := os.Args[1:]
	for _, next := range files {
		fmt.Println(next)
		byte_text, err := ioutil.ReadFile(next)
		if err != nil {
			fmt.Printf("File - %s - is not found", next)
			log.Fatal()
		}
		string_text := string(byte_text)
		fmt.Printf("%s", string_text)

	}
}
