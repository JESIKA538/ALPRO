package main

import (
	"fmt"
)

func main() {
	// Misalkan input bilangan x diberikan langsung di sini
	x := 253 // Gantilah nilai ini dengan angka yang ingin diuji (misalnya 253, 87, 999)

	// Memastikan input valid (x adalah bilangan bulat positif ≤ 999)
	if x < 0 || x > 999 {
		fmt.Println("Input tidak valid, masukkan bilangan bulat positif yang kurang dari atau sama dengan 999.")
		return
	}

	// Mendapatkan digit-digit dari bilangan x
	d1 := x / 100       // Ratusan
	d2 := (x / 10) % 10 // Puluhan
	d3 := x % 10        // Satuan

	// Menampilkan hasil
	fmt.Printf("Digit pertama (ratusan): %d\n", d1)
	fmt.Printf("Digit kedua (puluhan): %d\n", d2)
	fmt.Printf("Digit ketiga (satuan): %d\n", d3)
}
