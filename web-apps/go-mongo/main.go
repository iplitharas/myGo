package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-mongo/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello")

}

func getMongoClient() *mongo.Client {
	log.Println("Setting up MongoDB client..")
	// Connect to our local Mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}

func main() {
	mongoClient := getMongoClient()
	userController := controllers.NewUserController(mongoClient)
	router := httprouter.New()
	srv := http.Server{Addr: ":8081", Handler: router}
	log.Println("Listening at: http://localhost:8081")
	router.GET("/", index)
	router.GET("/user/:id", userController.GetUser)
	router.POST("/user/", userController.CreateUser)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error trying to start the server: %q", err)
	}
	defer func() {
		log.Println("Closing MongoDB client connection")
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
