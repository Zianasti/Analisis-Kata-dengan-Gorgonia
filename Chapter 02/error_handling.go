package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Pembagian oleh nol tidak diizinkan")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		return
	}
	fmt.Println("Hasil pembagian:", result)
}
