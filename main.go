package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	strcucts "github.com/teomandi/go-inaccess/pkg/structs"
)

func ptlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	period, tz := r.URL.Query().Get("period"), r.URL.Query().Get("tz")
	t1, t2 := r.URL.Query().Get("t1"), r.URL.Query().Get("t2")
	task, flag, msg := strcucts.MakeTask(period, tz, t1, t2)
	if !flag {
		errResponse := strcucts.ErrorResponse{
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
	addr := flag.String("addr", "localhost", "The address of the application")
	port := flag.String("port", "8000", "The port of the application")
	flag.Parse()

	r := mux.NewRouter()
	// r.Host(*addr + ":" + *port)
	r.HandleFunc("/ptlist", ptlist).Methods("GET")

	fmt.Printf("Starting server at %v:%v\n", *addr, *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
