package main

import "fmt"

func main() {
	map_1 := make(map[string]int)
	map_1["orange"] = 1
	map_1["blue"] = 2
	map_1["green"] = 3
	fmt.Println(map_1)

	map_2 := map[string]int{"joe": 123, "ian": 456}
	fmt.Println(map_2)

	for key, value := range map_2 {
		fmt.Println(key, value)
	}

	value, p := map_1["red"]
	if p {
		fmt.Println("red is in the map with value", value)
	} else {
		fmt.Println("red is not in the map", map_1)
	}

}
