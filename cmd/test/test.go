package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/teomandi/go-inaccess/pkg/structs"
)

var baseURL string = "http://localhost"
var port string = "8000" //should be parsed from env

func getPeriodicTimestamps(task structs.TestTask) bool {
	params := "period=" + url.QueryEscape(task.Period) + "&tz=" + url.QueryEscape(task.Timezone) + "&t1=" + url.QueryEscape(task.T1) + "&t2=" + url.QueryEscape(task.T2)
	path := fmt.Sprintf("%s:%s/ptlist?%s", baseURL, port, params)

	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if resp.StatusCode != 200 {
		return false
	}

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(body))
	return true
}

func main() {
	testTasks := []structs.TestTask{
		{Period: "1h", Timezone: "Europe/Athens", T1: "20210714T204603Z", T2: "20210715T123456Z"},
		{Period: "1d", Timezone: "Europe/Athens", T1: "20211010T204603Z", T2: "20211115T123456Z"},
		{Period: "1mo", Timezone: "Europe/Athens", T1: "20210214T204603Z", T2: "20211115T123456Z"},
		{Period: "1y", Timezone: "Europe/Athens", T1: "20180214T204603Z", T2: "20211115T123456Z"},
	}
	for i, testTask := range testTasks {
		if !getPeriodicTimestamps(testTask) {
			log.Fatal("Test ", i, " failed: "+testTask.ToString())
		}
	}
	log.Println("All tests completed with success!")
}
