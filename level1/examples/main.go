package main

import "fmt"

// menghitung luas lingkaran dan mengembalikan hasilnya
func hitungLuasLingkaran(jariJari float32) float32 {
	const phi float32 = 3.14
	return phi * jariJari * jariJari
}

// mencetak hasil perhitungan luas lingkaran
func cetakHasilPerhitungan(hasilHitung float32) {
	fmt.Printf("luas Lingkaran 3.2f", hasilHitung)
}

func main() {
	var jariJari float32 = 7.0
	LuasLingkaran := hitungLuasLingkaran(jariJari)
	cetakHasilPerhitungan(LuasLingkaran)

	var jariJari2 float32 = 10.0
	var jariJari3 float32 = 14.0
}
