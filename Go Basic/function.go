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

	fmt.Println(result)
}

func main() {
	fmt.Println("Hello my name is: ")

	getFullName()
}