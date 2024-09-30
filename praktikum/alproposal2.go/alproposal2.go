package main

import "fmt"

func main() {
	var nama, kelas string
	var nim string
	// fmt.Print("Masukkan nama, nim, kelas : ")
	// fmt.Scan(&nama, &kelas, &nim)
	fmt.Print("Masukkan nama")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan kelas anda")
	fmt.Scanln(&kelas)
	fmt.Print("Masukkan nim")
	fmt.Scanln(&nim)
	fmt.Println("Perkenalkan saya adalah ", nama, " salah satu mahasiswa Prodi S1-IF dari kelas ", kelas, " dengan NIM ", nim)

}
