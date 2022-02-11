package main

import "fmt"

func main() {
	name := "Dollong"

	switch name {
	case "Doltons":
		fmt.Println(true)
	default:
		fmt.Println("No one correct")
	}
	
	// short statement
	switch name == "Doltons" {
	case true:
		fmt.Println("Yes that true")
	case false:
		fmt.Println("No, thats false")
	}

	// switch without condition
	switch {
	case name == "Dolton":
		fmt.Println("True, variable name = Doltons")
	case name == "Dollong":
		fmt.Println("True, variable name = Dollong")
	}
}