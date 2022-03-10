package main

import "fmt"

func sumAll(text string, numbers ...int) {
	total := 0

	/**
	 * _ symbol in below is index. i dont need for this case
	 */
	for _, number := range numbers {
		total += number
	}

	fmt.Println(text, total)
}

func main() {
	sumAll("Total 10 + 10 + 10 =", 10, 10, 10)

	numbers := []int{10, 10, 10}

	// with slice argument
	/**
	 * ... symbol in below is used to give slice to variadic function
	 */
	sumAll("Total in slice of 10 + 10 + 10 =", numbers...)
}
