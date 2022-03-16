package main

import "fmt"

/**
 * Defer
 * defer is a built-in function from go that will run a function when a parent function finishes executing
 * no matter what happens, even if there is an error
 */

func logging() {
	fmt.Println("Done running function")
}

func runApp(isError bool) {
	defer logging()

	if isError {
		panic("App ERROR")
	}

	// this code will not to be running when isError == true
	fmt.Println("App is running")
}

func main() {
	runApp(true)
}
