package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
	"github.com/fatih/color"
	g "github.com/soniah/gosnmp"
)

func printValue(pdu g.SnmpPDU) error {
	fmt.Printf("%s = ", pdu.Name)

	switch pdu.Type {
	case g.OctetString:
		b := pdu.Value.([]byte)
		fmt.Printf("STRING: %s\n", string(b))
	default:
		fmt.Printf("TYPE %d: %d\n", pdu.Type, g.ToBigInt(pdu.Value))
	}
	return nil
}

func retPorts() {

	oid := ".1.3.6.1.4.1.43.10.26.1.1.1.5"

	err := g.Default.Walk(oid, printValue)
	if err != nil {
		fmt.Printf("Walk Error: %v\n", err)
		os.Exit(1)
	}
}

func setPort() {

	oid1 := ".1.3.6.1.4.1.43.10.26.1.1.1.5.1."

	var seg,port int
	fmt.Print("Enter segment: ")
	fmt.Scanf("%d", &seg)
	oid += "100" + strconv.Itoa(seg)

	fmt.Print("Enter port: ")
	fmt.Scanf("%d", &port)

	result, err := g.Default.Get([]string{oid1})
	if err != nil {
		fmt.Printf("Get() Error: %v\n", err)
		os.Exit(1)
	}

	for _, variable := range result.Variables {
		if variable.Type != g.OctetString {
			fmt.Printf("%d\n", g.ToBigInt(variable.Value))
		}
	}
}


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
	fmt.Println("5. Exit")
	fmt.Println("")
	fmt.Print("Choose an option: ")
	var i int
	fmt.Scanf("%d", &i)
	return i
}

func main() {
	ip := ""
	com := ""
	if len(os.Args) >= 3 {
		ip = os.Args[1]
		com = os.Args[2]
	}

	color.Red("Welcome to Cirquit")
	color.Yellow("Version 0.1")
	color.Green("Licensed under GNU Public License v3")
	fmt.Println("")
	reader := bufio.NewReader(os.Stdin)

	// Enter target IP
	if (ip == ""){
		fmt.Print("Enter target IP: ")
		ipN, _ := reader.ReadString('\n')
		ip = strings.TrimSuffix(ipN, "\n")
	}


	isValidIp := is_ipv4(ip)
	if !isValidIp {
		color.Red("The IP introduced is TRASH")
		os.Exit(3)
	}

	if (com == ""){
	// Enter community
		fmt.Print("Enter Community: ")
		comN, _ := reader.ReadString('\n')
		com = strings.TrimSuffix(comN, "\n")
	}

	g.Default.Target = ip
	g.Default.Community = com
	g.Default.Version = g.Version1
	g.Default.Timeout = time.Duration(10 * time.Second)


	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()



	exit := false
	for !exit {
		option := showMenu()

		switch option {
		case 1: retPorts()
		case 2: setPort()
		case 3:
		case 4:
		case 5:
			exit = true
		default:
			color.Red("Not a valid option\n\n")
		}
	}
}
