package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("testdata.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755) // Open file, append or create and give 755 perms
	if err != nil {
		return
	}
	defer file.Close()
	var user_data string
	fmt.Scanln(&user_data)
	if _, err := file.WriteString(user_data); err != nil {
		log.Println(err)
	}

}
