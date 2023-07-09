package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
)

func main() {
	g := gorgonia.NewGraph()
	x := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	y := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))

	// Menggunakan metode yang tidak didefinisikan pada objek
	z := x.Subtract(y) // Error: Undefined Method Error
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
    x := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
    y := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))
    // Menggunakan metode yang didefinisikan pada objek
    z := gorgonia.Must(gorgonia.Sub(x, y))
    fmt.Println(z.Value())
}
*/
