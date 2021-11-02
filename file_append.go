package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("testdata.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755) // Open file, append or create and give 755 perms
	if err != nil { // Error checking
		return
	}
	defer file.Close() // When the function is complete this will run

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	user_data := scanner.Text()

	if _, err := file.WriteString(user_data); err != nil {
		log.Println(err)
	}

}
