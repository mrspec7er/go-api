package main

import (
	"fmt"
	"log"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello There!")
}

func handleRequest()  {
	http.HandleFunc("/", greetings)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main()  {
	handleRequest()
}