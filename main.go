package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/fatih/color"
)

func showMenu() (int) {
	color.Magenta("MENU")
	fmt.Println("1. Retrive segments")
	fmt.Println("2. Modify port")
	fmt.Println("3. Retrieve IP layer details")
	fmt.Println("4. Change gateway")
	fmt.Println("")
	fmt.Print("Choose an option: ")
	var i int
	fmt.Scanf("%d", &i)
	return i
}

func main() {
	color.Red("Welcome to Cirquit")
	color.Yellow("Version 0.1")
	color.Green("Licensed under GNU Public License v3")
	fmt.Println("")

	// Enter target IP
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target IP: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	validEntry := false
	for !validEntry {
		option := showMenu()

		validEntry = true
		switch option {
		case 1:
		case 2:
		case 3:
		case 4:
		default:
			color.Red("Not a valid option\n\n")
			validEntry = false
		}
	}
}
