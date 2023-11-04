package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Response struct {
	Message string
}


func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Unable to load relevant env file for MongoDB cluster connection")
	}

	// Get environment variables
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASS")
	clusterAddress := os.Getenv("MONGODB_CLUSTER_ADDRESS")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + username + ":" + password + "@" + clusterAddress + "/testgolang?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	fmt.Println("mongodb+srv://" + username + ":" + password + "@" + clusterAddress + "/test?retryWrites=true&w=majority")
	client,err := mongo.Connect(context.TODO(),clientOptions)
	if(err != nil){ //error has occured during conn, does not check credentials at this point, separate check required, no err if wrong credentials
		fmt.Println("Unable to establish connection with mongo db")
	}	
	

	// Perform a ping to check if the credentials are correct and connection successfully established for operations
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		// Handle authentication or other specific connection errors
		fmt.Println("Authentication failed. Please chck credentials.")
		return
	}
	fmt.Println("Connection to MongoDB successful.")


	 // Create a new context
	 ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	 defer cancel()
 
	 // Create the database
	 err = client.Database("testgolang").CreateCollection(ctx, "golangdb1")
	 if err != nil {
		fmt.Println(err)

		return
	 }


	
	// var collection *mongo.Collection
	// collection = client.Database("testdb-go").Collection("messages")
	// fmt.Println(collection)
	
	

	// mux := http.NewServeMux()
	// mux.HandleFunc("/hello", withLogging(withCORS(helloHandler)))
	// fmt.Println("Server is running on http://localhost:8080")
	// http.ListenAndServe(":8080", mux)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
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
