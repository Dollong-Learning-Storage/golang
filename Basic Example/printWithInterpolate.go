package main

import "fmt"

func main() {
	mhsName := [...]string {
		"Doltons",
		"Edward",
		"Nico",
	}

	for i := 0; i < len(mhsName); i++ {
		// step 1
		// intToStr := fmt.Sprintf("%d", i);
		// fmt.Println("Name " + mhsName[i] + " is" + intToStr + " order") 

		// step 2
		fmt.Printf("Name %s is %v order \n", mhsName[i], i+1) 
	}
}