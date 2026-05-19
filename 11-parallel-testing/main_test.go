package main

import (
	"testing"
)

func TestTransfer_RaceCondition(t *testing.T) {
	// satu dompet yg akan di cek menggunakan 3 skenario secara bersamaan
	dompet := &DompetDigital{Saldo: 100}

	// tabel Data
	skenarioTes := []struct {
		namaKasus string
		tarik     int
	}{
		{"Req_1_tarik_50", 50},
		{"Req_2_tarik_50", 50},
		{"Req_3_tarik_50", 50},
	}

	// logic
	for _, skenario := range skenarioTes {
		// memecah tabel menjadi subtest yang bisa di isolasi
		t.Run(skenario.namaKasus, func(t *testing.T) {
			// menjalankan skenario secara serentak dgn afterburner
			t.Parallel()

			hasil := dompet.Transfer(skenario.tarik, "False_Rekening")

			// t.logf utk observasi
			// hanya muncul dgn flag -v
			t.Logf("[%s] Status: %s | Sisa Saldo di memory: %d", skenario.namaKasus, hasil, dompet.Saldo)
		})
	}
}

// run menggunakan go test -v dan
// go test -v -race untuk melihat detail perbedaannya
