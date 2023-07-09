package main

import "fmt"

func main() {
	nilai := 85
	if nilai >= 90 {
		fmt.Println("Nilai Anda sangat baik")
	} else if nilai >= 70 {
		fmt.Println("Nilai Anda memadai")
	} else {
		fmt.Println("Anda perlu meningkatkan nilai")
	}
}
