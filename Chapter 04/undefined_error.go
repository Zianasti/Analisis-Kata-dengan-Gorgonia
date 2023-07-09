package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
)

func main() {
	g := gorgonia.NewGraph()
	x := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	y := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))

	// Menggunakan metode yang tidak dikenali oleh Gorgonia
	z := x.Add(y) // perbaikan = z := gorgonia.Must(gorgonia.Add(x, y))

	// Menyusun dan menjalankan komputasi dengan Graph
	machine := gorgonia.NewTapeMachine(g)
	if err := machine.RunAll(); err != nil {
		fmt.Println("Error:", err)
	}
}
