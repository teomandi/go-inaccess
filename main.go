package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	strcucts "github.com/teomandi/go-inaccess/structs"
)

func ptlist(w http.ResponseWriter, r *http.Request) {
	arr := [3]string{"20211010T204603Z", "20211010T204603Z", "20211010T204603Z"}

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	period := params["period"]
	tz := params["tz"]
	t1 := params["t1"]
	t2 := params["t2"]

	task := strcucts.Task{Period: period, Timezone: "Europe/Athens", T1: t1, T2: t2}

	json.NewEncoder(w).Encode(arr)
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ptlist", ptlist).Methods("GET")

	fmt.Printf("Starting server at post 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

	// t1Str := "20211010T204603Z"
	// t2Str := "20211115T123456Z"
	// period := "1d"

	// t1, err := utils.ParseFormatedDate(t1Str)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// t2, err := utils.ParseFormatedDate(t2Str)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// task := strcucts.Task{Period: period, Timezone: "Europe/Athens", T1: t1, T2: t2}
	// timestamps := task.GetMatchingTimestamps()

	// for _, v := range timestamps {
	// 	fmt.Printf("%v\n", v)
	// }

}
