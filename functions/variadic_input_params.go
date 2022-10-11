package main

import "fmt"

func main() {
	in := make([]int, 0, 10)
	in = append(in, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("Input is:", in)
	out := addTo(10, in...)
	fmt.Println("Adding 10 to each element:", out)

}
func addTo(base int, values ...int) []int {
	out := make([]int, 0, len(values))
	for _, value := range values {
		out = append(out, base+value)
	}
	return out
}
