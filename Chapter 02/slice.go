package main

import "fmt"

func main() {
	// Membuat slice menggunakan literal
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers) // Output: [1 2 3 4 5]

	// Membuat slice dari array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println(slice) // Output: [2 3 4]
}
