package main

import "fmt"

func getFullName() (string, string) {
	return "Doltons", "Edward"
}

func getAlternateName() (string, string, string) {
	return "Edward", "Nico", "Pabiaran"
}

func main() {
	firstName, lastName := getFullName()
	alternateName, _, _ := getAlternateName()

	fmt.Println("My name is", firstName, lastName)
	fmt.Println("My alternate name", alternateName)
}
