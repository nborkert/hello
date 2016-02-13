package main

import (
	//	"fmt"
	"os"
	//	"net/url"
//	"time"
	"github.com/nborkert/getPage"
	"strconv"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//	fmt.Printf("Started...\n")
	//	fmt.Printf(os.Args[1])
	/*	msg := getPage.PrintMessage()
		more comments here
		fmt.Printf(msg)
		fmt.Printf("\n")
		msg2 := getPage.PrintEcho("ECHO TEST")
	 	fmt.Printf(msg2)
		fmt.Printf("\n")
	*/


	//masterRosterPage := "http://rankupdfsapi.com:3000/api/rosters/Baseball"

//	lineupPage := "http://rankupdfsapi.com:3000/api/teams/optimize/Baseball?userId=55b6fc5c4e1fc19066e8227d"
//	lineupPage := "https://104.236.233.0:3001/api/teams/optimize/Baseball?userId=55b6fc5c4e1fc19066e8227d"
//	lineupPage = "http://45.55.70.251:3000/api/teams/optimize/Baseball?userId=55b6fc5c4e1fc19066e8227d"
//	lineupPage = "http://45.55.70.251:3000"
//	lineupPage = "http://driven-striker-752.appspot.com/players"
//	lineupPage = "http://driven-striker-752.appspot.com"
//	lineupPage := "http://104.236.233.0:3000/"
//	lineupPage := "https://104.236.233.0:3001"
//	lineupPage := "http://rankupdfsapi.com:3000"

//RankUpTest2 below. Passes test. ALWAYS!
//	lineupPage := "http://159.203.70.245:3000"

//HAProxy below. Sometimes passes. Maybe because only one server on backend? Nope, because HAProxy thinks rapid serial requests are a DoS attack.
	lineupPage := "http://104.236.233.0:3000"

//HAProxy below.
//	lineupPage := "https://104.236.233.0:3001"


	max, _ := strconv.Atoi(os.Args[1])
	cresults := make(chan int)

//	for outer := 0; outer != 3; outer++ {

		for count := 0; count != max; count++ {
	//		go getPage.GetHTTPCode(masterRosterPage, count, cresults)
			go getPage.GetHTTPCode(lineupPage, count, cresults)
			//getPage.GetHTTPCodeNoChannel(lineupPage, count)
			//time.Sleep(time.Millisecond * 10)

		}


		for count := 0; count != max; count++ {
			<-cresults
		}


//	}

	/*
		body := getPage.GetHTTPCode(page)
		file, err := os.Create(u.Host)
		if err != nil {
			fmt.Printf("Couldn't open file")
		}
		defer file.Close()
		file.WriteString(body)
	*/
}
