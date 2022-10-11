/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create
a struct which contains the first and last names found in the file.
Each struct created will be added to a slice,
and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your
slice of structs and print the first and last names found in each struct.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

type PersonalInfo struct {
	fname string
	lname string
}

func (p *PersonalInfo) Print() {
	fmt.Printf("First name: %s -- Last name: %s\n", p.fname, p.lname)
}

func main() {
	filePath := FindFilePath()
	lines := OpenFile(filePath)
	persons := ParsePersons(lines)
	PrintPersons(persons)

}

func PrintPersons(persons []PersonalInfo) {
	fmt.Printf("Total #%d entries\n", len(persons))
	for _, person := range persons {
		person.Print()
	}
}

func ParsePersons(lines []string) []PersonalInfo {
	persons := make([]PersonalInfo, 0, 10)
	for _, line := range lines {
		firstName, LastName := SplitNames(line)
		person := PersonalInfo{
			fname: firstName,
			lname: LastName}
		persons = append(persons, person)
	}
	return persons
}

func SplitNames(line string) (string, string) {
	return strings.Split(line, " ")[0], strings.Split(line, " ")[1]
}

func FindFilePath() string {
	var fileName string
	fmt.Println("Enter your file name: ")
	_, err := fmt.Scan(&fileName)
	if err != nil {
		return ""
	}
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	fileName = strings.ToLower(fileName) + ".txt"
	return path.Join(pwd, fileName)

}

func OpenFile(filePath string) []string {
	lines := make([]string, 0, 10)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	return lines

}
