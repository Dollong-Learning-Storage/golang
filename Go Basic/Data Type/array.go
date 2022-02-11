package main

import "fmt"

func main() {
	// var names [3]string

	// names[0] = "Doltons"
	// names[1] = "Edward"
	// names[2] = "Nicholas"
	
	// var names = [3]string{
	// 	"Doltons",
	// 	"Edward",
	// 	"Nicholas",
	// }

	var names = [...]string {
		"Doltons",
		"Edward",
		"Nicholas",
	}

	fmt.Println(names)
}