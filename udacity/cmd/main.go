package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"math/rand"
	"net/http"
	"path"
	"strconv"
)

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Id        int    `json:"id"`
}

type CustomerResponse struct {
	Id int `json:"id"`
}

var Customers = make([]Customer, 0)
var ID = 0

func ListCustomers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	output, err := json.Marshal(&Customers)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Println("error during marshal:", err)
	}
	_, err = fmt.Fprintf(w, string(output))
	if err != nil {
		return
	}

}

func retrieveCustomer(customers []Customer, id int) (Customer, error) {
	for _, customer := range customers {
		if customer.Id == id {
			return customer, nil
		}
	}
	return Customer{}, errors.New("doesn't exist")
}

func deleteCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println("error is, ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	for _, customer := range Customers {
		if customer.Id == id {
		}
	}

}

func getCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println("error is, ", err)
		return
	}
	customer, err := retrieveCustomer(Customers, id)
	if err != nil {
		log.Printf("Cannot find customer with id %d\n", id)
		w.WriteHeader(http.StatusNotFound)
	}

	// serialize the customer
	output, err := json.Marshal(customer)
	_, err = w.Write(output)
	if err != nil {
		log.Println("error during write response", err)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Create new customer..")
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var customer Customer
	customer.Id = rand.Intn(100) + 1
	err := decoder.Decode(&customer)
	if err != nil {
		log.Printf("error during decoding reqeust: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Customers = append(Customers, customer)
	w.WriteHeader(http.StatusCreated)
	// serialize the new customer
	output, err := json.Marshal(CustomerResponse{Id: customer.Id})
	_, err = w.Write(output)
	if err != nil {
		log.Println("error during write response", err)
	}

}

func createSampleCustomers() {

	firstCustomer := Customer{
		FirstName: "John",
		LastName:  "Smith",
		Id:        1,
	}
	secondCustomer := Customer{FirstName: "Lara", LastName: "Croft", Id: 2}

	Customers = append(Customers, firstCustomer, secondCustomer)

}

func main() {
	createSampleCustomers()
	serverAddress := "http://localhost:8080"
	log.Printf("Serve server at: %s\n", serverAddress)
	mux := httprouter.New()
	mux.GET("/customers", ListCustomers)
	mux.POST("/customers", createCustomer)
	mux.GET("/customers/:id", getCustomer)
	mux.PUT("/customers/:id", getCustomer)

	server := http.Server{Addr: "localhost:8080", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		log.Println("error trying to server the server at: ", serverAddress, "is ", err)
		return
	}

}
