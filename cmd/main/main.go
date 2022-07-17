package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teomandi/go-inaccess/pkg/structs"
)

func ptlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	period, tz := r.URL.Query().Get("period"), r.URL.Query().Get("tz")
	t1, t2 := r.URL.Query().Get("t1"), r.URL.Query().Get("t2")
	task, flag, msg := structs.MakeTask(period, tz, t1, t2)
	if !flag {
		errResponse := structs.ErrorResponse{
			Status: "error",
			Desc:   msg,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}
	arr := task.GetMatchingTimestamps()
	json.NewEncoder(w).Encode(arr)
	return
}

func main() {
	// parse the command line arguments
	addr := flag.String("addr", "http://localhost", "The address of the application")
	port := flag.String("port", "8080", "The port of the application")
	flag.Parse()

	// initialize the app
	r := mux.NewRouter()
	r.HandleFunc("/ptlist", ptlist).Methods("GET")
	// r.Host(*addr + ":" + *port)

	fmt.Printf("Starting server at %v:%v\n", *addr, *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
