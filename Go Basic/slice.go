package main

import "fmt"

/*
	in slice we have pointer, length, capacity

	len(slice) -> to get length of slice
	cap(slice) -> to get capacity of slice

	make([]dataType, length, capacity)
	copy(toSlice, fromSlice)
*/

func main() {
	months := [...]string {
		"January",
		"February",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"July",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice1 := months[6:9]

	slice1[0] = "New"
	slice2 := append(slice1, "Desember Baru")
	fmt.Println(slice2)
	fmt.Println(slice1)

	dataMhs := make([]string, 2, 5)
	dataMhs[0] = "Dollong"
	dataMhs[1] = "Edw"

	fmt.Println(dataMhs)
	fmt.Println(cap(dataMhs))

	copySlice := make([]string, len(dataMhs), cap(dataMhs))

	copy(copySlice, dataMhs)

	copySlice[1] = "Edward"

	fmt.Println(copySlice)
}