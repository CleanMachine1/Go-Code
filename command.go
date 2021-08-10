package main

import (
	"fmt"
	"os/exec"
)

func main() {

	pw, err := exec.Command("pwd").Output()
	fmt.Println(err)
	fmt.Println(string(pw))
}
