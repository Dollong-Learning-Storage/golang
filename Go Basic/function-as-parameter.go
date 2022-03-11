package main

import "fmt"

func printText(name string, filter func(string) string) string {
	text := "My name is " + filter(name)
	return text
}

func filterText(name string) string {
	if name == "Anj" {
		return "***"
	}

	return name
}

func main() {
	fmt.Println(printText("Doltons", filterText))
	fmt.Println(printText("Anj", filterText))
}
