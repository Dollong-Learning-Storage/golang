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
		// continue statement
		if (key + 1) % 2 == 0 {
			continue
		}
		
		fmt.Println("index ", key + 1, " = ", name)

		// break statement
		// if key + 1 == 2 {
		// 	break
		// }
	}
}