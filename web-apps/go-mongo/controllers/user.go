package controllers

import (
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-mongo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type UserController struct {
	mongoClient *mongo.Client
}

// NewUserController UserController constructor
func NewUserController(mongoClient *mongo.Client) *UserController {
	return &UserController{mongoClient: mongoClient}
}

// GetUser View
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	user := models.User{}
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	objectId := bson.ObjectIdHex(id)
	filter := bson.M{"_id": objectId}
	uc.mongoClient.Connect(context.TODO())
	err := uc.mongoClient.Database("go-web-deb-db").Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("error during search of the the user: %q", err)
	}
	log.Println(user)
	err = json2.NewEncoder(w).Encode(user)
	if err != nil {
		errorMessage := fmt.Sprintf("error trying to serialize user with id:"+
			" %q is: %q", p.ByName("id"), err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
	}
}

// CreateUser view
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	user := models.User{}
	err := json2.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("error during deserialization:  %q\n", err)
	}
	// create bson ID
	user.Id = bson.NewObjectId()
	_, err = uc.mongoClient.Database("go-web-deb-db").Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		log.Panicf("error during storing user to the db: %q", err)
		return
	}
	err = json2.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("error during serialization %q: %q\n", user, err)
	}

}

// ListUsers view
func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	users := make([]models.User, 0, 10)
	filter := bson.M{}
	collection := uc.mongoClient.Database("go-web-deb-db").Collection("users")
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Panicf("error during search all the users: %q", err)
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Printf("error during executing the find-all cursor: %q", err)
	}
	userResponse, err := json2.Marshal(users)
	if err != nil {
		log.Panicf("error during serializasion of users %q", err)
	}
	fmt.Fprintf(w, string(userResponse))

}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	user := models.User{}
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	objectId := bson.ObjectIdHex(id)
	filter := bson.M{"_id": objectId}
	uc.mongoClient.Connect(context.TODO())
	collection := uc.mongoClient.Database("go-web-deb-db").Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("error is: %q", err)
		w.WriteHeader(http.StatusNotFound)
		return

	}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("error during deletion of the the user: %q", err)
	}
	log.Println(user)
	err = json2.NewEncoder(w).Encode(user)
	if err != nil {
		errorMessage := fmt.Sprintf("error trying to serialize user with id:"+
			" %q is: %q", p.ByName("id"), err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
	}

}
