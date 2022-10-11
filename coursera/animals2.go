package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Animal2 interface {
	Move() string
	Eat() string
	Speak() string
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c *Cow) Move() string {
	return c.locomotion
}

func (c *Cow) Eat() string {
	return c.food
}

func (c *Cow) Speak() string {
	return c.noise
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b *Bird) Eat() string {
	return b.food
}

func (b *Bird) Move() string {
	return b.locomotion
}

func (b *Bird) Speak() string {
	return b.noise
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s *Snake) Eat() string {
	return s.food
}

func (s *Snake) Move() string {
	return s.locomotion
}

func (s *Snake) Speak() string {
	return s.noise
}

func main() {
	fmt.Println("Available commands:\n" +
		"--> `newanimal` name type[cow|bird|snake]\n" +
		"--> `query` name info[eat|move|speak]")
	animals := make(map[string]Animal2)
	for {
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)
		inputLine, _ := reader.ReadString('\n')
		userInput := strings.Split(inputLine, " ")
		command, animalName, subCommand := userInput[0], strings.TrimSpace(userInput[1]), strings.TrimSpace(userInput[2])
		if command == "newanimal" {
			animals = createAnimals(animals, animalName, subCommand)
			printAvailableAnimals(animals)
		} else if command == "query" {
			printInfo(animals, animalName, subCommand)
		}

	}

}

func printAvailableAnimals(animals map[string]Animal2) {
	fmt.Printf("Total existing animals #%d\n", len(animals))
	animalNames := reflect.ValueOf(animals).MapKeys()
	fmt.Println("Available animals: ", animalNames)

}

func createAnimals(animals map[string]Animal2, animalName string, animalType string) map[string]Animal2 {
	switch animalType {
	case "cow":
		{
			cow := Cow{
				food:       "grass",
				locomotion: "walk",
				noise:      "moo",
			}
			animals[animalName] = &cow
			fmt.Println("Created, ", animalName)
			return animals
		}
	case "bird":
		{
			bird := Bird{
				food:       "worms",
				locomotion: "fly",
				noise:      "peep",
			}
			animals[animalName] = &bird
			fmt.Println("Created, ", animalName)
			return animals
		}
	case "snake":
		{
			snake := Snake{
				food:       "mice",
				locomotion: "slither",
				noise:      "hsss",
			}
			animals[animalName] = &snake
			fmt.Println("Created, ", animalName)
			return animals
		}
	default:
		fmt.Println("Available types are: `cow`, `bird` `snake`")
		fmt.Println("Requested type: ", animalType)

	}
	return animals
}

func printInfo(animals map[string]Animal2, animalName string, option string) {
	animal, ok := animals[animalName]
	if ok {
		switch option {
		case "eat":
			{
				fmt.Println(animal.Eat())
				return
			}
		case "move":
			{
				fmt.Println(animal.Move())
				return
			}
		case "speak":
			{
				fmt.Println(animal.Speak())
				return
			}
		default:
			fmt.Println("Available options are: `eat`, `move` `speak`")
			fmt.Println("Requested option: ", option)
			return

		}
	}
	fmt.Printf("Can't find the requested animal with name: %s\n", animalName)
	printAvailableAnimals(animals)

}
