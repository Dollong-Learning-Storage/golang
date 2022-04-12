package main

import "fmt"

/**
 * Struct
 * Struct is a data template or prototype data
 * Struct is a data template that is used to combine zero or more data types
 * Structs are usually data representations of the applications that we create
 * The data from the struct is collected in a field
 * Simply put, a struct is a collection of fields
 */

/**
 * We can't use the structure right away
 * But we can create data/objects from the structs we create
 */
type Mhs struct {
	Name, Address string
	Age, Nim      int
}

func main() {
	// var doltons Mhs
	// doltons.Name = "Doltons Edward"
	// doltons.Address = "JL. Kompleks"
	// doltons.Age = 19
	// doltons.Nim = 10229

	// struct literals
	// step 1 (Best practice)
	// doltons := Mhs{
	// 	Name:    "Doltons Edward",
	// 	Address: "JL. Musyawarah",
	// 	Age:     19,
	// 	Nim:     21038,
	// }

	// step 2
	doltons := Mhs{"Doltons Edward", "JL. Musyawarah", 19, 23923}

	fmt.Println(doltons.Name)
}
