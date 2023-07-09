package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
)

func main() {
	g := gorgonia.NewGraph()

	// Deklarasi tensor x dan y dengan ukuran yang tidak sesuai
	x := gorgonia.NewTensor(g, gorgonia.Float64, 2, gorgonia.WithShape(2, 3))
	y := gorgonia.NewTensor(g, gorgonia.Float64, 3, gorgonia.WithShape(3, 2))

	// Operasi matematika antara x dan y
	z, err := gorgonia.Mul(x, y)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Buat VM untuk mengeksekusi graph
	machine := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(x, y))

	// Eksekusi graph
	if err := machine.RunAll(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(z.Value())
}
