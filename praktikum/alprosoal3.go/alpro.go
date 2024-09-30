package main

import "fmt"

func main() {
	var r, luaslingkaran float64
	fmt.Print("Masukkan nilai r: ")
	fmt.Scan(&r)
	luaslingkaran = 3.14 * (r * r)
	fmt.Println(luaslingkaran)
}
