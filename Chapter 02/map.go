package main

import "fmt"

func main() {
	// Membuat map dengan tipe data string sebagai kunci dan int sebagai nilai
	student := make(map[string]int)
	student["John"] = 90
	student["Alice"] = 95
	fmt.Println(student)
	// Output: map[John:90 Alice:95]
}
