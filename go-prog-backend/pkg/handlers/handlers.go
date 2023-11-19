package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChengCK18/go-proj-backend/pkg/database"
)

type Response struct {
	Message string
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	var response Response

	switch r.Method {
	case http.MethodGet:
		response = Response{Message: "Hello from Houston, hearing you loud and clear (GET)"}
	case http.MethodPost:
		HelloHandlerPost(w, r)
	case http.MethodPut:
		response = Response{Message: "Hello from Houston, processing your PUT request"}
	case http.MethodDelete:
		response = Response{Message: "Hello from Houston, handling your DELETE request"}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// response := Response{Message: "Hello from Houston, hearing you loud and clear "+r.Method}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HelloHandlerPost(w http.ResponseWriter, r *http.Request) {
	var userInput database.SampleData

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid JSON data, missing name and/or age field", http.StatusBadRequest)
		return
	}

	err2 := database.InsertIntoMongoDB(userInput)
	if err2 != nil {
		http.Error(w, "Failed to insert data into MongoDB", http.StatusInternalServerError)
		return
	}

	// Respond with a success message or appropriate response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Data inserted successfully")
}
