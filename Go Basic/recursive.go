package main

import "fmt"

// step 1: with for loop
// func factorial(num int) int {
// 	var result int

// 	for i := num; i > 0; i-- {
// 		result *= i
// 	}

// 	return result
// }

func factorial(num int) int {
	if num == 1 {
		return 1
	}

	// instead of using if else, i use return early concept
	// ref: https://dev.to/jpswade/return-early-12o5
	return num * factorial(num-1)
}

func main() {
	fiveFactor := factorial(5)
	fmt.Println(fiveFactor)
}
