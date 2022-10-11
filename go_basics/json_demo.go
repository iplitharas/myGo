package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

func main() {

	dog := Animal{Age: 2, Name: "Rex"}
	var dogBack Animal
	fmt.Println("I am a: ", dog)
	bytes, err := json.Marshal(dog)
	if err != nil {
		return
	}
	fmt.Println("Dog as json, array of bytes", bytes)
	err = json.Unmarshal(bytes, &dogBack)
	if err != nil {
		fmt.Println("Error during Unmarshal: ", err)
		return
	}
	fmt.Println("Dog again", dogBack)
}
