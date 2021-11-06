package main

import (
	"fmt"
)

func main() {
	var data int
	fmt.Scanf("%d", &data)
	fmt.Println(data)
	if data%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
