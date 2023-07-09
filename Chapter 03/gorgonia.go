package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
)

func main() {
	// Buat objek graf komputasi
	g := gorgonia.NewGraph()
	// Buat tensor input
	x := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	// Berikan nilai pada tensor input
	xVal := 10.0
	gorgonia.Let(x, xVal)
	// Definisikan operasi komputasi
	y := gorgonia.Must(gorgonia.Add(x, gorgonia.NewConstant(5.0)))
	// Buat VM (Virtual Machine) untuk menjalankan graf komputasi
	machine := gorgonia.NewTapeMachine(g)
	// Jalankan graf komputasi
	if err := machine.RunAll(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Ambil hasil komputasi dari tensor output
	result := y.Value().Data().(float64)
	// Cetak hasil
	fmt.Println("Hasil komputasi:", result)
}
