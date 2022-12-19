package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-mongo/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello")

}

// func getMongoClient create and return a connection to the
// local mongodb
func getMongoClient() *mongo.Client {
	log.Println("Setting up MongoDB client..")
	// Connect to our local Mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	log.Println("Successfully connected and pinged MongoDB")
	return client
}

func main() {
	client := getMongoClient()
	defer func() {
		log.Println("Disconnecting from mongoDB")
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println("error during mongo disconnect: %q", err)
		}
	}()

	userController := controllers.NewUserController(client)
	router := httprouter.New()
	srv := http.Server{Addr: ":8081", Handler: router}
	log.Println("Listening at: http://localhost:8081")
	router.GET("/", index)
	router.GET("/user/:id", userController.GetUser)
	router.DELETE("/user/:id", userController.DeleteUser)
	router.POST("/user/", userController.CreateUser)
	router.GET("/users/", userController.ListUsers)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to start the server: %q", err)
	}
	defer func() {
		log.Println("Closing MongoDB client connection")
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
