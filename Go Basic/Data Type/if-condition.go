package main

import "fmt"

func main() {
	x := 5

	if x > 5 {
		fmt.Println("Thats true")
	} else if x < 5 {
		fmt.Println("Almost right")
	} else {
		fmt.Println("No, thats wrong")
	}

	// short statement 
	if xIsBiggerThan5 := x > 5; xIsBiggerThan5 {
		fmt.Println("Nop")
	}
}