package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byte_text, err := ioutil.ReadFile("testdata.txt") // Gets the entire file
	if err != nil {
		fmt.Println(err)
	}
	// byte_text is currently in a byte format, unreadable
	// fmt.Println(byte_text)
	string_text := string(byte_text)
	fmt.Println(string_text)
}
