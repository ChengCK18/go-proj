package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChengCK18/go-proj-backend/pkg/database"
	"github.com/ChengCK18/go-proj-backend/pkg/model"
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
	var userInput model.SampleData

	// verify if incoming data can be decoded into JSON
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		fmt.Println("Invalid JSON data")
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Check all relevant fields are populated
	if userInput.Name == "" || userInput.Age <= 0 {
		fmt.Println("Invalid user data, missing name and/or age field")
		http.Error(w, "Invalid user data, missing name and/or age field", http.StatusBadRequest)
		return
	}

	// Insert parsed user data into mongoDB
	err2 := database.InsertIntoMongoDB(userInput)
	if err2 != nil {
		fmt.Println("Failed to insert data into MongoDB")
		http.Error(w, "Failed to insert data into MongoDB", http.StatusInternalServerError)
		return
	}
	
	
	// Respond with a success message or appropriate response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Data inserted successfully")
}
