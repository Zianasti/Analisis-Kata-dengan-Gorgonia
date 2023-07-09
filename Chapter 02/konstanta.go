package main

import "fmt"

func main() {
	const phi = 3.14
	const jariJariLingkaran = 5
	luasLingkaran := phi * jariJariLingkaran * jariJariLingkaran
	fmt.Println(luasLingkaran) // Output: 78.5

	const umurMinimal = 18
	var umur = 20
	bolehDewasa := umur >= umurMinimal
	fmt.Println(bolehDewasa) // Output: true

	const salam = "Halo, "
	const nama = "John"
	pesanSalam := salam + nama
	fmt.Println(pesanSalam) // Output: Halo, John
}
