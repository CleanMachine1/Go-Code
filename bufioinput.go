package main

// Made since input is seperated with spaces, which is annoying in entering user data
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Make scanner
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line) // Print what is recorded
}
