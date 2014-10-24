package main

import (
	"fmt"
	"os"
	"net/url"
	"github.com/nborkert/getPage"
)

func main() {
	fmt.Printf("Started...\n")
	fmt.Printf(os.Args[1])
/*	msg := getPage.PrintMessage()
	fmt.Printf(msg)
	fmt.Printf("\n")
	msg2 := getPage.PrintEcho("ECHO TEST")
 	fmt.Printf(msg2)
	fmt.Printf("\n")
*/
	//page := "http://www.example.com"
	page := os.Args[1]
	u, err := url.Parse(page)
	if err != nil {
		panic(err)
	}

	body := getPage.GetContent(page)
	file, err := os.Create(u.Host)
	if err != nil {
		fmt.Printf("Couldn't open file")
	}
	defer file.Close()
	file.WriteString(body)
}
