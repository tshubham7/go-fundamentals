package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/json", myJSONHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// myJSONHandler
func myJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// we will call thrid party api, get the data and display on our api
	resp, err := http.Get("https://api.github.com/repositories/19438/commits")
	if err != nil {
		w.Write([]byte("unable to fetch data from api"))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte("unable to read data from api"))
	}
	w.Write(body)
}
