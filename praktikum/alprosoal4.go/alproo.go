package main

import "fmt"

func main() {
	var fahrenheit float64
	// Meminta input suhu dalam Fahrenheit dan pengguna
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)
	// Menghitung suhu dalam Celcius menggunakan rumus C=(F-32)*5/9
	celcius := (fahrenheit - 32) * 5 / 9
	// Menampilkan hasil suhu dalam Celcius
	fmt.Print("Suhu dalam celcius:%.lf\n", celcius)
}
