package main

import "fmt"

func main() {
	defer fmt.Println("Hello world")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
