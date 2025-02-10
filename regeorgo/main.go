package main

import (
	"github.com/bhagyashriw777/regeorgo"
	"net/http"
)

func main() {
	// initialize regeorgo
	gh := &regeorgo.GeorgHandler{LogLevel: 0}
	gh.InitHandler()

	// use it as standard handler for http
	http.HandleFunc("/regeorgo", gh.RegHandler)
	http.ListenAndServe(":8111", nil)
}
