package main

import "fmt"

func main() {
	//	slices is a window of an underline array
	// pointer -> indicates the start
	// len
	// capacity
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(arr)
	slice_1 := arr[1:3]
	slice_2 := arr[2:5]
	fmt.Println(slice_1)
	fmt.Println(slice_2)
	slice_2 = append(slice_2, "yolo")
	fmt.Println(slice_2)
	slice_3 := make([]string, 1, 1)
	fmt.Println("Slice 3 is", slice_3)
	fmt.Println("Capacity is", cap(slice_3))
	fmt.Println("Length is", len(slice_3))
	fmt.Println("Add two elements to this slice")
	slice_3 = append(slice_3, "add something")
	slice_3 = append(slice_3, "add again")
	fmt.Println("Slice 3 is", slice_3)
	fmt.Println("Capacity is", cap(slice_3))
	fmt.Println("Length is", len(slice_3))
	// difference between array creation
	// we don't specify the size or ....
	slice_4 := []int{1, 2, 3}
	fmt.Println("Creating slice literal", slice_4)

}
