package main

import "fmt"

type Hello struct {
	X string
	Y int
}

func main() {
	v := Hello{"Hi", 2} // x = hi Y = 2
	fmt.Printf(v.X)     // "Hi"
	v.X = "How are you?"
	fmt.Printf(v.X) // How are you?
}
