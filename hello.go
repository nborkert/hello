package main

import (
	"github.com/nborkert/getPage"
	"os"
	"runtime"
	"strconv"
)

//Command-line argument is number of concurrent requests to send to the target
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	targetPage := "http://52.36.17.139:80"

	max, _ := strconv.Atoi(os.Args[1])
	cresults := make(chan int)

	for outer := 0; outer != 5; outer++ {

		for count := 0; count != max; count++ {
			go getPage.GetHTTPCode(targetPage, count, cresults)
		}

		//Now insure we have all the results before beginning the next test round
		for count := 0; count != max; count++ {
			<-cresults
		}

	}

}
