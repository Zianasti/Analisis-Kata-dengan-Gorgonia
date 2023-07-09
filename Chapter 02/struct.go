package main

import "fmt"

// Definisikan struct Student
type Student struct {
	Name  string
	Grade int
	Age   int
}

func main() {
	// Inisialisasi struct Student
	student1 := Student{
		Name:  "John",
		Grade: 10,
		Age:   15,
	}

	// Mengakses nilai-nilai dalam struct
	fmt.Println("Name:", student1.Name)   // Output: Name: John
	fmt.Println("Grade:", student1.Grade) // Output: Grade: 10
	fmt.Println("Age:", student1.Age)     // Output: Age: 15
}
