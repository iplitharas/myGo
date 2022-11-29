package main

import (
	"errors"
	"fmt"
)

// ResourceError custom error
// where the resource can be db or some other resource
// code can be an integer
type ResourceError struct {
	Resource string
	code     int
}

func (re ResourceError) Error() string {
	return fmt.Sprintf("%s %d", re.Resource, re.code)
}

func (re ResourceError) Is(target error) bool {
	fmt.Println("-----Entering `Is` with target----", target)
	if other, ok := target.(ResourceError); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.code == other.code
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}

func createTwoResourceErrors() (error, error) {
	firstError := ResourceError{
		Resource: "Database",
		code:     1,
	}
	secondError := ResourceError{
		Resource: "Database",
		code:     2,
	}
	return firstError, secondError
}

func main() {
	fmt.Println("Custom error checking with the `errors.Is`")
	firstError, _ := createTwoResourceErrors()
	if errors.Is(firstError, ResourceError{
		Resource: "Database",
		code:     -1,
	}) {
		fmt.Println("The database is broken!")
	}
}
