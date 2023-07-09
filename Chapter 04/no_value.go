package main

func main() {
	x := 5
	y := 10 // Variabel y dideklarasikan tetapi tidak digunakan
	_ = x   // Penggunaan variabel x
}

/*perbaikan :
package main
func main() {
    x := 5
    _ = x // Variabel x digunakan, sedangkan variabel y dihapus karena tidak diperlukan
}*/
