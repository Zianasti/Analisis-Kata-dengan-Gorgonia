package main

import "fmt"

type PersegiPanjang struct {
	panjang, lebar float64
}

func (p PersegiPanjang) Luas() float64 {
	return p.panjang * p.lebar
}

func main() {
	p := PersegiPanjang{panjang: 4, lebar: 5}
	luas := p.Luas()
	fmt.Println(luas) // Output: 20
}
