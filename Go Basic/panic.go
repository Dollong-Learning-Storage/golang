package main

import "fmt"

/**
 * Panic
 * panic is the function we use to stop the program
 * The panic function is usually called when an error occurs while the program is running
 * when panic is called, the program stops but defer will still run
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
