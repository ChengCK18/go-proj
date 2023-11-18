package handlers

import (
	"net/http"
	"encoding/json"

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
			response = Response{Message: "Hello from Houston, receiving your POST request"}
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
