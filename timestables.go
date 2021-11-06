package main

import "fmt"

func main() {

	var number, i int // The number being multiplied by
	fmt.Scanf("%d", &number)
	for i = 1; i <= 10; i++ { // Starts at x1, goes to 10, repeats the timing ten times
		output := i * number
		fmt.Printf("%d x %d = %d\n", number, i, output) // Print the total for each run

	}

}
