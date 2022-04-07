package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	crwl "webcrawler/src/crawler"

	"github.com/gorilla/mux"
)

// output struct
type Apiresponse struct {
	Value []crwl.Responsebody `json:"result"`
	Err   string              `json:"error"`
}

//input struct
type inputUrl struct {
	Urllist []string `json:"urls"`
}

//take input and purduce output for web crawling
func crawlapi(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly provide data in correct format")
	}

	var newEvent inputUrl
	json.Unmarshal(reqBody, &newEvent)
	w.WriteHeader(http.StatusCreated)

	var resp Apiresponse
	cr, err := crwl.Webcrawler(newEvent.Urllist)
	if err != nil {
		resp = Apiresponse{
			Value: cr,
			Err:   err.Error(),
		}
	} else {
		resp = Apiresponse{
			Value: cr,
			Err:   "",
		}
	}

	json.NewEncoder(w).Encode(resp)
}

//handleRequests: handles the request for rest api call
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/crawl", crawlapi).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Hello there starting web crawler")
	handleRequests()
}
