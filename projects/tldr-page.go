package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/* To add to file:
title1 \n
desc
link \n
up to 8 examples
*/
/* Todo:

[]Save as command-name.md

[]implement a system for amount of command examples, either ask before or just when reaching end of each
[]include an 8 cap on the amounts

[x]Use the correct syntax, eg. periods and backticks

[x]Create and append, deleting the old one with user permission
*/

// Page creator for TLDR-pages, simplifying the syntax

func main() {

	fmt.Println("Enter the name of the program/command:")
	scanner := bufio.NewScanner(os.Stdin) // Create a text scanner
	scanner.Scan()                        // Scan for the title
	title1 := scanner.Text()              // Collect this run of the scan and save
	title1 = strings.TrimSuffix(title1, " ") // Removes commonly applied extra space when entering values
	pagename := strings.ReplaceAll(title1, " ", "-") + ".md" // for creating the file nam
	if _, err := os.Stat(pagename); err == nil {
		fmt.Printf("file %q already exists, overwrite it? (y/N)", pagename)
		scanner.Scan()
		choice := scanner.Text()
		if choice == "y" || choice == "yes" || choice == "Yes" {
			os.Remove(pagename)
			os.Create(pagename)
		} else {
			fmt.Println("Exiting")
			os.Exit(1)
		}

	}
	title1 = "# " + title1 // For when writing to page

	fmt.Println("Enter a description for the program/command:")
	scanner.Scan()
	desc := scanner.Text()
	desc = "- " + desc + "."

	fmt.Println("Enter a more information link:")
	scanner.Scan()
	link := scanner.Text()
	link = "More information: " + link + "."

	file, err := os.OpenFile(pagename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Close the file at the final command
	file.WriteString(title1 + "\n")
	file.WriteString("\n" + desc)
	file.WriteString("\n" + link + "\n")
	var i int
	for i = 0; i <= 8; i++ { // commands part of the page
		fmt.Println("Enter a description for a command example:")
		scanner.Scan()
		command_desc := scanner.Text()
		command_desc = "- " + command_desc + ":"
		file.WriteString("\n" + command_desc + "\n")
		fmt.Println("Enter a corresponding command for the given description:")
		scanner.Scan()
		command := scanner.Text()
		command = "`" + command + "`"
		file.WriteString("\n" + command + "\n")
		if i < 8 {
			fmt.Println("Do you want to add another command to the page? (y/N)")
			scanner.Scan()
			another := scanner.Text()
			if another == "y" || another == "yes" || another == "Yes" {
				fmt.Println("Moving to next command!")
			} else {
				i = 9
			}
		}
	}
}
