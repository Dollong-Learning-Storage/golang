package main

import "fmt"

/*
	uint8 as byte
	rune as int32
	int as Minimal int32
	uint as minimal uint32
*/

func main() {
	var (
		x byte = 10
		y byte = 15
	)
	
	sum := x + y

	var name = "Doltons Edward"
	const constantName = "Dollong"

	// build in function go
	nameLength := len(name)
	lastChatOfName := name[0]
	
	fmt.Println(nameLength)
	fmt.Println(lastChatOfName)
	fmt.Println(sum)
	fmt.Println(constantName)
}