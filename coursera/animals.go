package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mammals interface {
	Eat() string
	Move() string
	Speak() string
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func main() {

	for {
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)
		inputLine, _ := reader.ReadString('\n')
		userInput := strings.Split(inputLine, " ")
		animal, info := userInput[0], userInput[1]
		info = strings.TrimSpace(info)
		if animal == "cow" {
			cow := Animal{
				food:       "grass",
				locomotion: "walk",
				noise:      "moo",
			}
			MammalInfo(&cow, info)

		}
		if animal == "bird" {
			bird := Animal{
				food:       "worms",
				locomotion: "fly",
				noise:      "peep",
			}
			MammalInfo(&bird, info)
		}
		if animal == "snake" {
			snake := Animal{
				food:       "mice",
				locomotion: "slither",
				noise:      "hsss",
			}
			MammalInfo(&snake, info)
		}
	}

}

func MammalInfo(m Mammals, info string) {
	switch info {
	case "eat":
		{
			fmt.Println(m.Eat())
		}
	case "move":
		{
			fmt.Println(m.Move())
		}
	case "speak":
		{
			fmt.Println(m.Speak())
		}
	}

}
