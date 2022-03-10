package main

import "fmt"

/**
 * alternate we can use
 * func getFullName() (firstName string, middleName string, lastName string)
 */
func getFullName() (firstName, middleName, lastName string) {
	firstName = "Doltons"
	middleName = "Edward"
	lastName = "Nicholas"

	return
}

func main() {
	/**
	 * we can choose any name for this variable, example
	 * a, b, c := getFullName()
	 */
	firstName, _, lastName := getFullName()

	fmt.Println("My full name is", firstName, lastName)
}
