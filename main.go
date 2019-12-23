package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
	"github.com/fatih/color"
	g "github.com/soniah/gosnmp"
)

func is_ipv4(host string) bool {
	parts := strings.Split(host, ".")
	if len(parts) < 4 {
		return false
	}
	for _,x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

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
	ipN, _ := reader.ReadString('\n')
	ip := strings.TrimSuffix(ipN, "\n")
	isValidIp := is_ipv4(ip)
	if !isValidIp {
		color.Red("The IP introduced is TRASH")
		os.Exit(3)
	}

	g.Default.Target = ip
	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

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
