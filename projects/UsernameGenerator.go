package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomint1 := randInt(0, 1868)
	randomint2 := randInt(0, 6800)
	randomint3 := randInt(0, 9999)

	content1, err := ioutil.ReadFile("adjective.txt")
	if err != nil {
		fmt.Println(err)
	}

	adjectives := strings.Split(string(content1), "\n")

	content2, err := ioutil.ReadFile("noun.txt")
	if err != nil {
		fmt.Println(err)
	}

	nouns := strings.Split(string(content2), "\n")

	username := adjectives[randomint1] + nouns[randomint2] + strconv.Itoa(randomint3)

	fmt.Printf("Your generated username: %s", username)
}
