package main

import "fmt"

func getData() [3]string {
	fullName := [...]string{
		"Doltons",
		"Edward",
		"Nicholas",
	}

	return fullName
}

func getFullName() {
	fullName := getData()

	var result string
	for i := 0; i < len(fullName); i++ {
		result += fullName[i] + " "
	}

	fmt.Printf("%s, %d length of name", result, len(result))
}

func sayHelloTo(firstName string, lastName string) string {
	fullName := firstName + " " + lastName
	return "Hello " + fullName
}

func main() {
	fmt.Println("Hello my name is: ")

	getFullName()
	fmt.Printf("\n")
	fmt.Println(sayHelloTo("Eka", "Dwi"))
}
