package main

import (
	"crypto/tls"
	"github.com/nborkert/getPage"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	max, _ := strconv.Atoi(os.Args[1])
	delay, _ := strconv.Atoi(os.Args[2])
	targetPage := os.Args[3]

	//	c := make (chan int)
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}

	for count := 0; count != max; count++ {
		go getPage.GetHTTPCodeNoChannel(targetPage, count, client)
		//		go getPage.GetHTTPCode(targetPage, count, c, client)
	}

	/*
		for count := 0; count != max; count++ {
			<-c;
		}
	*/

	time.Sleep(time.Second * time.Duration(delay))

}
