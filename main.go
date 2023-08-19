package main

import "fmt"

func main() {
	var a, b, pilihan, hasil int

	fmt.Println("Kalkulator Sederhana")
	fmt.Println("1. Pertambahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")
	fmt.Println("Masukkan pilihan Anda: ")
	fmt.Scanln(&pilihan)

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&b)

	switch pilihan {
  case 1:
		hasil = a + b
	case 2:
		hasil = a - b
  case 3:
		hasil = a * b
	}

	fmt.Println("Hasilnya adalah:", hasil)
}
