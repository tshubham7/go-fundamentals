package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/djson", myJSONHandler)
	addr := ":8080"
	fmt.Println("listening on ", addr)
	http.ListenAndServe(addr, nil)
}

// myJSONHandler
// this will only fetch data and write to the http response
func myJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// we will call thrid party api, get the data and display on our api
	resp, err := http.Get("https://api.github.com/repositories/19438/commits")
	if err != nil {
		handlerError(w, r, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handlerError(w, r, err)
		return
	}
	w.Write(body)
}

func handlerError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s", "message": "request failed"}`, err.Error())))
}
