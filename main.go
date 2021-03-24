package main

import (
	"fmt"
	"log"
	"net/http"
)

func electricity(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "There is electricity on AWS EC2")
}

func handleRequests() {
	http.HandleFunc("/", electricity)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func main() {
	handleRequests()
}
