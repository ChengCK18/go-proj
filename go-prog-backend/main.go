package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", withLogging(withCORS(helloHandler)))
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello from Houston, hearing you loud and clear"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling request => ", r.URL.Path)
		next(w, r)
	}
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next(w, r)
	}
}
