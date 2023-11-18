package main

import (

	"fmt"
	"net/http"
	"github.com/ChengCK18/go-proj-backend/pkg/handlers"
)



func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", withLogging(withCORS(handlers.HelloHandler)))
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)


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
