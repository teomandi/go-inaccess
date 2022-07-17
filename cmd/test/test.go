package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/teomandi/go-inaccess/pkg/structs"
)

var addr *string
var port *string

// requests from the app the given test task and checks the results
func getPeriodicTimestamps(task structs.TestTask) bool {
	params := "period=" + url.QueryEscape(task.Period) + "&tz=" + url.QueryEscape(task.Timezone) + "&t1=" + url.QueryEscape(task.T1) + "&t2=" + url.QueryEscape(task.T2)
	path := fmt.Sprintf("%s:%s/ptlist?%s", *addr, *port, params)
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if resp.StatusCode != 200 {
		return false
	}
	return true
}

func main() {
	// parse the command line arguments
	addr = flag.String("addr", "http://localhost", "The address of the application")
	port = flag.String("port", "8080", "The port of the application")
	flag.Parse()
	log.Printf("Testing at %s:%s\n", *addr, *port)

	// init the test-tasks
	testTasks := []structs.TestTask{
		{Period: "1h", Timezone: "Europe/Athens", T1: "20210714T204603Z", T2: "20210715T123456Z", Result: true},
		{Period: "1d", Timezone: "Europe/Athens", T1: "20211010T204603Z", T2: "20211115T123456Z", Result: true},
		{Period: "1mo", Timezone: "Europe/Athens", T1: "20210214T204603Z", T2: "20211115T123456Z", Result: true},
		{Period: "1y", Timezone: "Europe/Athens", T1: "20180214T204603Z", T2: "20211115T123456Z", Result: true},
		{Period: "1w", Timezone: "Europe/Athens", T1: "20180214T204603Z", T2: "20211115T123456Z", Result: false},
	}
	//testing the test-tasks
	for i, testTask := range testTasks {
		if getPeriodicTimestamps(testTask) != testTask.Result {
			log.Fatal("Test ", i, " failed: "+testTask.ToString())
		}
	}
	// if we are here all the task succeed
	log.Println("All tests completed with success!")
}
