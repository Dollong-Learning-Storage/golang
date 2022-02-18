package main

import "fmt"

/*
	in the golang there is only one feature for looping that is for loop
*/

func main() {
	fullName := []string {
		"Doltons",
		"Edward",
		"Nicholas",
		"Pabiaran",
	}

	for i := 0; i < len(fullName); i++ {
		fmt.Println(fullName[i])
	}

	// for loop with range
	for key, name := range fullName {
		fmt.Println("index ", key, " = ", name)
	}
}