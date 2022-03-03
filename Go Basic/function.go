package main

import "fmt"

func getFullName() {
	fullName := [...]string {
		"Doltons",
		"Edward",
		"Nicholas",
	}

	var result string
	for i := 0; i < len(fullName); i++ {
		result += fullName[i] + " "
	}

	fmt.Printf("%s, %d length of name", result , len(result))
}

func main() {
	fmt.Println("Hello my name is: ")

	getFullName()
}