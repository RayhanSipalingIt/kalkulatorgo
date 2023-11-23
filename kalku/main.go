// main.go

package main

import (
	"errors"
	"fmt"
)

// Fungsi-fungsi operasi matematika

func Penjumlahan(a, b int) int {
	return a + b
}

func Pengurangan(a, b int) int {
	return a - b
}

func Perkalian(a, b int) int {
	return a * b
}

func Pembagian(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("pembagian oleh nol")
	}
	return a / b, nil
}

func main() {
	var a, b int

	// Input dua angka
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&b)

	// Menampilkan menu pilihan operasi
	fmt.Println("Pilih operasi:")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")

	var choice int
	fmt.Print("Masukkan nomor operasi: ")
	fmt.Scan(&choice)

	// Menentukan operasi berdasarkan pilihan pengguna
	switch choice {
	case 1:
		result := Penjumlahan(a, b)
		fmt.Printf("Hasil penjumlahan: %d\n", result)
	case 2:
		result := Pengurangan(a, b)
		fmt.Printf("Hasil pengurangan: %d\n", result)
	case 3:
		result := Perkalian(a, b)
		fmt.Printf("Hasil perkalian: %d\n", result)
	case 4:
		result, err := Pembagian(a, b)
		//jika memasukkan angka 0 maka error
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Hasil pembagian: %d\n", result)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
