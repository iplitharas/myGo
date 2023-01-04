// Demo of using Generics to be able to implement
// functions where they can accept any Type
// Without generics we would need to implement
// custom functions for each type

package main

import "fmt"

// Stringer custom interface
type Stringer = interface {
	String() string
}

// String is a wrapper on string which implements the Stringer interface
type String string

func (s String) String() string {
	return string(s)
}

// Integer  is wrapper on int which implements the Stringer interface
type Integer int

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

// Student simple demo struct
type Student struct {
	firstName string
	lastName  string
	age       int
}

func (s Student) String() string {
	return fmt.Sprintf("%s %s %d", s.firstName, s.lastName, s.age)
}

// addStudent accepts any type of Stringer
func addStudent[T Stringer](students []T, student T) []T {
	return append(students, student)
}

func main() {
	// add some pure strings to the list of the students
	studentsStrings := make([]String, 0)
	studentsStrings = addStudent[String](studentsStrings, "Ioannis")
	studentsStrings = addStudent[String](studentsStrings, "Michael")
	fmt.Println("Students (strings) are: ", studentsStrings)

	// add some Students
	students := make([]Student, 0)
	students = addStudent[Student](students, Student{
		firstName: "Ioannis",
		lastName:  "Plitharas",
		age:       32,
	})
	fmt.Println("Students are: ", students)
	// add some integers
	studentsInts := make([]Integer, 0)
	studentsInts = addStudent[Integer](studentsInts, 1)
	studentsInts = addStudent[Integer](studentsInts, 2)
	fmt.Println("students (ints) are: ", studentsInts)

}
