package main

import "fmt"

/**
 * Closure
 * This closure feature allows us to interact with whatever is around us,
 * but only for the scope of that function or once above it
 */

func main() {
	count := 0
	fmt.Println("start count is", count)

	increment := func() {
		fmt.Printf("Increment count from %d to %d \n", count, count+1)
		count++
	}

	increment()
	increment()
	fmt.Println("final count is", count)
}
