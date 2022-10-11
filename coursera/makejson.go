package main

import (
	"encoding/json"
	"fmt"
)

/*
* Write a program which prompts the user to first enter a name, and then enter an address.
* Your program should create a map and add the name and address to the map using the keys “name” and “address”,
* respectively. Your program should use Marshal() to create a JSON object from the map,
*and then your program should print the JSON object.
 */
func main() {
	var firstName string
	var address string
	fmt.Println("Enter your firstname: ")
	_, err := fmt.Scan(&firstName)
	if err != nil {
		return
	}
	fmt.Println("Enter your address: ")
	_, err = fmt.Scan(&address)
	if err != nil {
		return
	}
	fmt.Println("Your name:", firstName, "and your address:", address)
	data := map[string]string{"name": firstName, "address": address}
	barry, _ := json.Marshal(data)
	fmt.Println("The json as a binary array:", barry)
	fmt.Println("as json format:", string(barry))

	dataBack := make(map[string]string)
	err = json.Unmarshal(barry, &dataBack)
	if err != nil {
		fmt.Println("Error during loading the json: ", err)
		return
	}
	fmt.Println("Load from json:", dataBack)

}
