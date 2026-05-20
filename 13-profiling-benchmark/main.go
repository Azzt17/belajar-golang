package main

import "strings"

// INI BURUK, karena memicu Memory Bloat dan menyiksa Garbage Collector
func GabungStringBuruk(jumlah int) string {
	var hasil string
	for i := 0; i < jumlah; i++ {
		hasil += "Golang" // Alokasi RAM baru terus terjadi
	}
	return hasil
}

// BAGUS, Menggunakan arsitektur Buffer untuk meminimalkan alokasi RAM
func GabungStringBagus(jumlah int) string {
	var builder strings.Builder
	for i := 0; i < jumlah; i++ {
		builder.WriteString("Golang")
	}
	return builder.String()
}

func main() {
	// Program utama dibiarkan kosong karena kita fokus pada Benchmark
}
