package main

import (
	"fmt"
	"github.com/nborkert/getPage"
)

func main() {
	fmt.Printf("Started...\n")
	msg := getPage.PrintMessage()
	fmt.Printf(msg)
	fmt.Printf("\n")
	msg2 := getPage.PrintEcho("ECHO TEST")
 	fmt.Printf(msg2)
	fmt.Printf("\n")
	body := getPage.GetContent("http://www.example.com")
	fmt.Printf(body)
	
}
