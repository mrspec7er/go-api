package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello There!")
}

func stockMarket(w http.ResponseWriter, r *http.Request) {
	stock := []string{"BBCA", "BBNI", "GOTO"}
	json.NewEncoder(w).Encode(stock)
}

func getStockData(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	json.NewEncoder(w).Encode(id)
}

// HandleFunc
func handleRequestWithHandleFunc()  {
	http.HandleFunc("/", greetings)
	http.HandleFunc("/stocks", stockMarket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//ServeMux
func handleRequestWithServeMux()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greetings)
	mux.HandleFunc("/stocks", stockMarket)
	mux.HandleFunc("/stocks/", getStockData)

	server := http.Server{
		Addr: ":8000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func main()  {
	// handleRequestWithHandleFunc()
	handleRequestWithServeMux()
}

