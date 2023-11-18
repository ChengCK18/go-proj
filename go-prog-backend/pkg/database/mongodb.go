// package database

// import (
// 	"context"
// 	"log"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var client *mongo.Client

// func init() {

// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("Unable to load relevant env file for MongoDB cluster connection")
// 	}

// 	// Get environment variables
// 	username := os.Getenv("MONGODB_USERNAME")
// 	password := os.Getenv("MONGODB_PASS")
// 	clusterAddress := os.Getenv("MONGODB_CLUSTER_ADDRESS")


// 	// Set up MongoDB connection
// 	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
// 	clientOptions := options.Client().ApplyURI("mongodb+srv://" + username + ":" + password + "@" + clusterAddress + "/testgolang?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
// 	var err error
// 	client, err = mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}


// 	// Perform a ping to check if the credentials are correct and connection successfully established for operations
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		// Handle authentication or other specific connection errors
// 		fmt.Println("Authentication failed. Please chck credentials.")
// 		return
// 	}
// 	fmt.Println("Connection to MongoDB successful.")

// 	defer func() {
// 		if err = client.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()
// }


// func InsertIntoMongoDB(data userData) error {
// 	collection := client.Database("testgolang").Collection("golangdb1")

// 	// Insert data into MongoDB
// 	_, err := collection.InsertOne(context.TODO(), data)
// 	return err
// }
