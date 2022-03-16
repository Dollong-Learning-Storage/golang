package main

import "fmt"

/**
 * Recover
 * recover is the function we call to catch the panic data
 * by recovering the panic process will stop then the program will continue to run
 */

func endApp() {
	errMessage := recover()

	if errMessage != nil {
		fmt.Println("Error with message", errMessage)
	}

	fmt.Println("SD")
}

func runApp(isError bool) {
	// we need to use defer to be able to use resolve
	defer endApp()

	if isError {
		panic("App ERROR")
	}

	fmt.Println("App is running")
}

func main() {
	runApp(true)
}
