package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
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
		HelloHandlerGet(w, r)
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

func HelloHandlerGet(w http.ResponseWriter, r *http.Request){
	var userInput model.SampleData
	var name string //optional parameter
	fmt.Println(r.Body)


	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil && err != io.EOF { // if it's not empty body related
		fmt.Println("[GET]Invalid JSON data")
		http.Error(w, "[GET]Invalid JSON data", http.StatusBadRequest)
		return
		
	}else{ // no error, there are field(s) in the r.Body
		
		if(userInput.Name != ""){
			name = userInput.Name
		}
	}
	

	results, _ := database.GetFromMongoDB(name)
	for _, result := range results {
        fmt.Printf("Name: %s, Age: %d\n", result.Name, result.Age)
    }
}

func HelloHandlerPost(w http.ResponseWriter, r *http.Request) {
	var userInput model.SampleData

	// verify if incoming data can be decoded into JSON
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		fmt.Println("[POST]Invalid JSON data")
		http.Error(w, "[POST]Invalid JSON data", http.StatusBadRequest)
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
