// di skenario ini kita membangun sistem live viewer counter, normalnya menggunakan count++ saja sudah cukup,
// namun sebenarnya mekanisme di tingkat CPU untuk count++ adalah melakukan Read dari RAM ke CPU,
// kemudian menambahkan satu angka, lalu melakukan Write kembali ke RAM
// jika dua atau lebih goroutine melakukan proses ini bersamaan, mereka akan saling menimpa hasil
// ini di sebut Lost Update Vulnerability
// menggunakan mutex untuk menambal kerentanan ini sangat overkill karena menambah kompleksitas proses dan kode
// karena itu untuk perhitungan numerik sederhana cukup menggunakan sync/atomic
// ini adlaah operasi lock-free yang memberikan instruksi langsung ke CPU untuk menyelesaikan 1 siklus proses
// yang tak bisa di interupsi oleh apapun (atomik)

package main

import (
	"fmt"
	"sync"
	"sync/atomic" // tambahkan ini untuk operasi tingkat hardware
)

func main() {
	fmt.Println("--SIMULASI LIVE VIEWER: ATOMIC COUNTER--")

	var totalViewer int32 = 0
	var wg sync.WaitGroup

	// ada 10.000 penonton yang masuk bersamaan
	jumlahPenontonMasuk := 10000

	for i := 0; i < jumlahPenontonMasuk; i++ {
		wg.Add(1)

		// goroutine yg mencatat penonton masuk
		go func() {
			defer wg.Done()

			// kita jadikan proses Read, Proses, Write dari counter
			// menjadi operasi tunggal (atomik) yg tak bisa di interpusi
			// oleh goroutine lain ketika berjalan
			atomic.AddInt32(&totalViewer, 1)
		}()
	}

	wg.Wait()

	// hasilnya selalu akurat 10.000, dan lebih cepat daripada Mutex
	fmt.Printf("[AUDIT] Sistem mencatat %d penonton masuk.\n", totalViewer)
	fmt.Printf("[AUDIT] Jumlah tiket yang hilang/tidak tercatat: %d\n", jumlahPenontonMasuk-int(totalViewer))
}

// gunakan race detector untuk run kode ini (go run -race)
