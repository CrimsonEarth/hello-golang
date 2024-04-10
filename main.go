package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	displayMenu()
}

func displayMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var (
			input string
			err   error
		)
		fmt.Println("\nWelcome to the Hello-GoLang CLI")
		fmt.Println("1. Print Hello")
		fmt.Println("2. Exit")

		fmt.Println("\nEnter your choice:")
		input, err = reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		// trim whitespace
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("Hello World!")
		case "2":
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice! Try again!")
		}
	}
}
