package main

import "fmt"

type BlackList func(string) bool

func register(name string, blackList BlackList) string {
	if blackList(name) {
		return "You are not supposed to be here"
	} else {
		return "Welcome " + name
	}
}

// func blackListAdmin(name string) bool {
//     return name == "admin"
// }

func main() {
	// step 1
	blackListAdmin := func(name string) bool {
		return name == "admin"
	}

	fmt.Println(register("admin", blackListAdmin))
	fmt.Println(register("Doltons", blackListAdmin))

	// step 2
	fmt.Println(register("root", func(name string) bool {
		return name == "root"
	}))
}
