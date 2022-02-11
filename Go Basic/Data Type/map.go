package main

import "fmt"

/*
	map[keyDataType]valueDataType
	map[string]string
	map[string]string {
		"key": "value"
	}

	delete(map, key)
*/

func main() {
	// map based on key-value
	person := map[string]string {
		"name": "Dollong",
		"address": "Maros",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["title"])
	
	delete(person, "title")
	fmt.Println(person)
}