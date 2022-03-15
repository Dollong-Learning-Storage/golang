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

func runApp(num int) int {
	/**
		 * Defer must always be at the beginning of a function. so that it can run as expected even in an error
	     * Defer function will always be executed last
	*/
	defer logging() // this will be running

	fmt.Println("App is running")
	result := 10 / num
	fmt.Printf("Result of 10 / %d is %d", num, result)

	return result
}

func main() {
	runApp(0)
}
