// di skenario ini kita membangun sistem live viewer counter, normalnya menggunakan count++ saja sudah cukup,
// namun sebenarnya mekanisme di tingkat CPU untuk count++ adalah melakukan Read dari RAM ke CPU,
// kemudian menambahkan satu angka, lalu melakukan Write kembali ke RAM
// jika dua atau lebih goroutine melakukan proses ini bersamaan, mereka akan saling menimpa hasil
// ini di sebut Lost Update Vulnerability

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("--SIMULASI LIVE VIEWER: BAD COUNTER--")

	var totalViewer int32 = 0
	var wg sync.WaitGroup

	// ada 10.000 penonton yang masuk bersamaan
	jumlahPenontonMasuk := 10000

	for i := 0; i < jumlahPenontonMasuk; i++ {
		wg.Add(1)

		// goroutine yg mencatat penonton masuk
		go func() {
			defer wg.Done()

			// karena ini bukan operasi tunggal, ada kemungkinan goroutine akan saling timpa
			// ketika CPU melakukan context switch di tengah2 prosesnya
			totalViewer++
		}()
	}

	wg.Wait()

	// hasilnya menjadi acak dan tak pernah mencapai 10.000
	fmt.Printf("[AUDIT] Sistem mencatat %d penonton masuk.\n", totalViewer)
	fmt.Printf("[AUDIT] Jumlah tiket yang hilang/tidak tercatat: %d\n", jumlahPenontonMasuk-int(totalViewer))
}

// gunakan race detector untuk run kode ini (go run -race)
