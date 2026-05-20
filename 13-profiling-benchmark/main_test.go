package main

import "testing"

// Penamaan wajib diawali dengan Benchmark, parameter b *testing.B
func BenchmarkGabungStringBuruk(b *testing.B) {
	// b.N akan diisi otomatis oleh Go (bisa 1, 1000, atau 1.000.000)
	for i := 0; i < b.N; i++ {
		// Menggabungkan kata "Golang" sebanyak 10.000 kali per eksekusi
		GabungStringBuruk(10000)
	}
}

func BenchmarkGabungStringBagus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GabungStringBagus(10000)
	}
}

// jalankan menggunakan 'go test -bench=. -benchmem'
// untuk melakukan bechmark test

// gunakan 'go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out'
// untuk mengekstark file profiling

// gunakan 'go tool pprof mem.out'
// untuk visualisator dan "top10"
// untuk menampilkan 10 fungsi paling memakan RAM di visualisator tadi
