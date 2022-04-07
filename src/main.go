package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	crwl "webcrawler/src/crawler"
)

type Apiresponse struct {
	Value []crwl.Responsebody `json:"result"`
	Err   error               `json:"error"`
}

func homepage(w http.ResponseWriter, r *http.Request) {
	cr, err := crwl.Webcrawler()
	resp := Apiresponse{
		Value: cr,
		Err:   err,
	}
	json.NewEncoder(w).Encode(resp)
}

func handleRequests() {
	http.HandleFunc("/api/crawl", homepage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	fmt.Println("hello world")
	handleRequests()
}
