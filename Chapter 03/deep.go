package main

import (
	"fmt"

	gorgonia "gorgonia.org/gorgonia"
)

func main() {
	// Contoh penggunaan Gorgonia untuk tugas deep learning
	fmt.Println("Contoh penggunaan Gorgonia untuk tugas deep learning")

	// Definisikan objek graf komputasi
	g := gorgonia.NewGraph()

	// Definisikan tensor input
	x := gorgonia.NewMatrix(g, gorgonia.Float64, gorgonia.WithShape(1, 2), gorgonia.WithName("x"))
	gorgonia.WithInit(gorgonia.GlorotU(1))(x)

	// Definisikan parameter model
	w := gorgonia.NewMatrix(g, gorgonia.Float64, gorgonia.WithShape(2, 1), gorgonia.WithName("w"))
	gorgonia.WithInit(gorgonia.GlorotU(1))(w)
	b := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("b"))
	gorgonia.WithInit(gorgonia.Zeroes())(b)

	// Hitung prediksi
	pred, _ := gorgonia.Add(gorgonia.Must(gorgonia.Mul(x, w)), b)

	// Buat Virtual Machine dan jalankan graf komputasi
	machine := gorgonia.NewTapeMachine(g)
	defer machine.Close()
	if err := machine.RunAll(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Ambil nilai hasil prediksi dari tensor pred
	result := pred.Value().Data().([]float64)

	// Cetak hasil prediksi
	fmt.Println("Hasil prediksi:", result)
}
