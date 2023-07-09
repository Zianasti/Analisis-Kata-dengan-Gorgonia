package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
)

func main() {
	g := gorgonia.NewGraph()
	x := gorgonia.NewScalar(g, gorgonia.Float64)
	y := gorgonia.NewScalar(g, gorgonia.Float64)
	z := gorgonia.Must(gorgonia.Add(x, y))
	machine := gorgonia.NewTapeMachine(g)
	defer machine.Close()
	xVal := 2.0
	yVal := 3.0
	gorgonia.Let(x, xVal)
	gorgonia.Let(y, yVal)
	if err := machine.RunAll(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	result := z.Value().Data().(float64)
	fmt.Println("Hasil penjumlahan:", result)
}
