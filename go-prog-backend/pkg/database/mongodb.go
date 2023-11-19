package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SampleData struct {
	Name string // need to be caps to be exported
	Age  int    // need to be caps to be exported
}

var client *mongo.Client

func init() {

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Unable to load relevant env file for MongoDB cluster connection")
		fmt.Println(err)
		return
	}

	// Get environment variables
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASS")
	clusterAddress := os.Getenv("MONGODB_CLUSTER_ADDRESS")

	// Set up MongoDB connection
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + username + ":" + password + "@" + clusterAddress + "/testgolang?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Perform a ping to check if the credentials are correct and connection successfully established for operations
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		// Handle authentication or other specific connection errors
		fmt.Println("Authentication failed. Please chck credentials.")
		return
	}
	fmt.Println("Connection to MongoDB successful.")

	// if put defer here, it will cause the client to disconnect upon completion of init func
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
}

func InsertIntoMongoDB(data SampleData) error {
	collection := client.Database("testgolang").Collection("golangdb1")

	// Insert data into MongoDB
	_, err := collection.InsertOne(context.TODO(), data)
	fmt.Println("geeee", err)
	return err
}
