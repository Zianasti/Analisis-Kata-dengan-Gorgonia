package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
)

func main() {
	g := gorgonia.NewGraph()
	x := gorgonia.NewTensor(g, gorgonia.Float32, 2, gorgonia.WithShape(2, 2), gorgonia.WithName("x"))
	y := gorgonia.NewTensor(g, gorgonia.Int, 2, gorgonia.WithShape(2, 2), gorgonia.WithName("y"))

	// Operasi penjumlahan antara dua tensor dengan tipe data yang tidak kompatibel
	z, err := gorgonia.Add(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(z.Value())
}

/* perbaikan :
package main

import (
    "fmt"

    "gorgonia.org/gorgonia"
)

func main() {
    g := gorgonia.NewGraph()
    x := gorgonia.NewTensor(g, gorgonia.Float32, 2, gorgonia.WithShape(2, 2), gorgonia.WithName("x"))
    y := gorgonia.NewTensor(g, gorgonia.Float32, 2, gorgonia.WithShape(2, 2), gorgonia.WithName("y"))
    z, err := gorgonia.Add(x, y)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(z.Value())
}
*/
