package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	g "github.com/soniah/gosnmp"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

func setPortCheck() {
    
    var seg, port int
    
    var aux []string
    
    
    
    if _, err := os.Stat("ports.cfg"); err == nil {
        fmt.Println("FILE DETECTED")
        //File exists
        file, _ := os.Open("ports.cfg")
        fscanner := bufio.NewScanner(file)
        for fscanner.Scan() {
            if (fscanner.Text()[0] != ';'){
                aux = strings.Split(fscanner.Text(), ":") 
                
                
                seg, _ = strconv.Atoi(aux[1])
                port, _ = strconv.Atoi(aux[0])
                
                setPort(seg,port)
            }
            
        }
        

    } else {
        // Obtenemos primero el valor asociado al segmento 2:
        fmt.Println("IGNORE FILE")
        fmt.Print("Enter port: ")
        fmt.Scanf("%d", &port)

        fmt.Print("Enter segment: ")
        fmt.Scanf("%d", &seg)
        
        setPort(seg, port)
            
    }
}

func setPort(seg int, port int){
    
    oid := ".1.3.6.1.4.1.43.10.26.1.1.1.5.1."
    
	
	oid1 := oid + "100" + strconv.Itoa(seg)

	oid2 := oid + strconv.Itoa(port)

	result, err := g.Default.Get([]string{oid1})
	if err != nil {
		fmt.Printf("Get() Error: %v\n", err)
		os.Exit(1)
	}

	var segInt interface{}
	var seg2Int interface{}

	for _, variable := range result.Variables {
		if variable.Type != g.OctetString {
			//fmt.Printf("%d\n", g.ToBigInt(variable.Value))
			segInt = (variable.Value)
		}
	}

	// Lo establecemos para el puerto port:
	pdu := g.SnmpPDU{
		Name:  oid2,
		Type:  g.Integer,
		Value: segInt,
	}

	_, err = g.Default.Set([]g.SnmpPDU{pdu})

	result, err = g.Default.Get([]string{oid2})
	for _, variable := range result.Variables {
		if variable.Type != g.OctetString {
			//fmt.Printf("%d\n", g.ToBigInt(variable.Value))
			seg2Int = (variable.Value)
		}

	}
	if seg2Int == segInt {
		c := color.New(color.FgGreen)
		c.Println("The port " +  strconv.Itoa(port) + " has been changed" )
	} else {
		c := color.New(color.FgRed)
		c.Println("Error in port " +  strconv.Itoa(port))
	}
}

func getIP() {

	oidConf := ".1.3.6.1.4.1.43.10.27.1.1.1.15.1"

	oidIP := ".1.3.6.1.4.1.43.10.28.1.1.2."
	oidMask := ".1.3.6.1.4.1.43.10.28.1.1.3."
	oidGate := ".1.3.6.1.4.1.43.10.28.1.1.4."
	var resInt interface{}

	//Comprobar los valores de direccionamiento IP:
	result, err := g.Default.Get([]string{oidConf})
	if err != nil {
		fmt.Printf("Get() Error: %v\n", err)
		os.Exit(1)
	}
	for _, variable := range result.Variables {
		if variable.Type != g.OctetString {
			//fmt.Printf("%d\n", g.ToBigInt(variable.Value))
			resInt = (variable.Value)
		}

	}
	// Para comprobar el valor de la dirección IP del hub:

	oidIP += fmt.Sprintf("%v", resInt)

	result, err = g.Default.Get([]string{oidIP})
	if err != nil {
		fmt.Printf("Get() Error: %v\n", err)
		os.Exit(1)
	}
	for _, variable := range result.Variables {
		fmt.Print("IP Address: ")
		c := color.New(color.FgYellow)
		c.Println(variable.Value)
	}

	// Para comprobar el valor de la máscara de la dirección IP del hub:

	oidMask += fmt.Sprintf("%v", resInt)

	result, err = g.Default.Get([]string{oidMask})
	if err != nil {
		fmt.Printf("Get() Error: %v\n", err)
		os.Exit(1)
	}
	for _, variable := range result.Variables {
		fmt.Print("Subnet Mask: ")
		c := color.New(color.FgCyan)
		c.Println(variable.Value)
	}

	// Para comprobar el valor de la dirección IP del router por defecto:

	oidGate += fmt.Sprintf("%v", resInt)

	result, err = g.Default.Get([]string{oidGate})
	if err != nil {
		fmt.Printf("Get() Error: %v\n", err)
		os.Exit(1)
	}
	for _, variable := range result.Variables {
		fmt.Print("Default Router: ")
		c := color.New(color.FgBlue)
		c.Println(variable.Value)
	}

}

func setIP() {
	oid := ".1.3.6.1.4.1.43.10.28.1.1.4."
	oidConf := ".1.3.6.1.4.1.43.10.27.1.1.1.15.1"

	//Comprobar los valores de direccionamiento IP:
	var resInt interface{}
	result, err := g.Default.Get([]string{oidConf})
	if err != nil {
		fmt.Printf("Get() Error: %v\n", err)
		os.Exit(1)
	}
	for _, variable := range result.Variables {
		if variable.Type != g.OctetString {
			resInt = (variable.Value)
		}

	}
	// Para comprobar el valor de la dirección IP del hub:

	oid += fmt.Sprintf("%v", resInt)

	// Obtenemos primero el valor asociado al segmento 2:

	var ip string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter gateway IP: ")
	ipN, _ := reader.ReadString('\n')
	ip = strings.TrimSuffix(ipN, "\n")

	isValidIp := is_ipv4(ip)
	if !isValidIp {
		color.Red("The IP introduced is TRASH")
		os.Exit(3)
	}

	// Lo establecemos para el puerto port:
	pdu := g.SnmpPDU{
		Name:  oid,
		Type:  g.IPAddress,
		Value: ip,
	}

	fmt.Println(oid)

	g.Default.Set([]g.SnmpPDU{pdu})

}

func is_ipv4(host string) bool {
	parts := strings.Split(host, ".")
	if len(parts) < 4 {
		return false
	}
	for _, x := range parts {
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

func showMenu() int {
	color.Magenta("-------MENU-------")
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
	color.Yellow("Version 0.2")
	color.Green("Licensed under GNU Public License v3")
	fmt.Println("")
	reader := bufio.NewReader(os.Stdin)

	// Enter target IP
	if ip == "" {
		fmt.Print("Enter target IP: ")
		ipN, _ := reader.ReadString('\n')
		ip = strings.TrimSuffix(ipN, "\n")
	}

	isValidIp := is_ipv4(ip)
	if !isValidIp {
		color.Red("The IP introduced is TRASH")
		os.Exit(3)
	}

	if com == "" {
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
		case 1:
			retPorts()
		case 2:
			setPortCheck()
		case 3:
			getIP()
		case 4:
			setIP()
		case 5:
			exit = true
		default:
			color.Red("Not a valid option\n\n")
		}
	}
}
