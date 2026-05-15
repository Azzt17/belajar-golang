// ini contoh bagaimana race condition bisa terjadi jika sebuah variabel tidak di batasi akses
// read dan writingnya ketika banyak requet terjadi bersamaan
package main

import (
	"fmt"
	"sync"
	"time"
)

// struct yg rentan (tidak memiliki mekanisme penguncian)
type DompetDigital struct {
	Saldo int
}

// funsgi ini memiliki celah Time-of-Check Time-of-Use (TOCTOU)
func (d *DompetDigital) TarikUang(jumlah int, namaUser string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Time-of-Check
	if d.Saldo >= jumlah {
		fmt.Printf("[%s] Mengecek saldi: %d. Cukup! Memproses penarikan..\n", namaUser, d.Saldo)
		// simulasi delay jaringan (celah ini bisa di manfaatkan oleh hacker)
		// ketika goroutine ini sleep, go scheduler akan memasukkan goroutin lain
		// yang juga lolos pengecekan if di atas
		time.Sleep(1 * time.Millisecond)

		// Time-of-Use
		d.Saldo -= jumlah
		fmt.Printf("[%s] Berhasil menarik %d. Sisa saldo: %d\n", namaUser, jumlah, d.Saldo)
	} else {
		fmt.Printf("[%s] GAGAL menarik %d. Saldo tidak cukup, Sisa saldo: %d\n", namaUser, jumlah, d.Saldo)
	}
}

func main() {
	fmt.Println("--SIMULASI PERETASAN TOCTOU--")

	// Korban memiliki saldo awal 100.000
	dompet := &DompetDigital{Saldo: 100000}
	var wg sync.WaitGroup

	// hacker melakukan request bersamaan secara pararel
	wg.Add(2)
	go dompet.TarikUang(80000, "Req_1_Hacker", &wg)
	go dompet.TarikUang(80000, "Req_2_Hacker", &wg)

	wg.Wait()

	// Saldo menjadi minus!
	fmt.Printf("[AUDIT] Saldo Akhir di Database: %d\n", dompet.Saldo)
}

// RUN MENGGUNAKAN "go run -race bad_toctou.go" untuk mendeteksi race condition
