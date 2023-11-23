// mathops.go

package mathops

import "errors"

// Penjumlahan mengembalikan hasil penjumlahan dua bilangan.
func Penjumlahan(a, b float64) float64 {
	return a + b
}

// Pengurangan mengembalikan hasil pengurangan dua bilangan.
func Pengurangan(a, b float64) float64 {
	return a - b
}

// Perkalian mengembalikan hasil perkalian dua bilangan.
func Perkalian(a, b float64) float64 {
	return a * b
}

// Pembagian mengembalikan hasil pembagian a oleh b.
// Mengembalikan error jika b sama dengan nol.
func Pembagian(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("pembagian oleh nol")
	}
	return a / b, nil
}
