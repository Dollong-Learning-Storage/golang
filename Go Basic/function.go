package main

import "fmt"

func getFullName() {
	fullName := [...]string{
		"Doltons",
		"Edward",
		"Nicholas",
	}

	var result string
	for i := 0; i < len(fullName); i++ {
		result += fullName[i] + " "
	}

	fmt.Printf("%s, %d length of name", result, len(result))
}

func sayHelloTo(firstName string, lastName string) {
	fullName := firstName + " " + lastName
	fmt.Println("Hello", fullName)
}

func main() {
	fmt.Println("Hello my name is: ")

	getFullName()
	fmt.Printf("\n")
	sayHelloTo("Eka", "Dwi")
}
