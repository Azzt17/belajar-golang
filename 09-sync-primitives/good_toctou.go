// untuk case yg sama, kita menjadikan fase check dan use sebagai satu kesatuan yg
// tidak bisa di interupsi ketika di proses, menggunakan sync.Mutex

package main

import (
	"fmt"
	"sync"
	"time"
)

// struct yg aman (menggunakan gembok)
type DompetDigital2 struct {
	Saldo int
	Mu    sync.Mutex // tambahkan mutex agar menempel dgn datanya
}

func (d *DompetDigital2) TarikUang(jumlah int, namaUser string, wg *sync.WaitGroup) {
	defer wg.Done()

	// kita mengunci aksesnya sebelum melakukan pengecekan
	// jika ada goroutine lain yg memanggil Lock, maka akan di tahan (antri)
	// sampai goroutine ini memanggil Unlock()
	d.Mu.Lock()

	// pastikan aksesnya akan di buka setelah fungsi selesai, bahkan jika terjadi panic
	defer d.Mu.Unlock()

	// Time-of-Check
	if d.Saldo >= jumlah {
		fmt.Printf("[%s] Mengecek saldi: %d. Cukup! Memproses penarikan..\n", namaUser, d.Saldo)
		time.Sleep(1 * time.Millisecond)

		// Time-of-Use
		d.Saldo -= jumlah
		fmt.Printf("[%s] Berhasil menarik %d. Sisa saldo: %d\n", namaUser, jumlah, d.Saldo)
	} else {
		fmt.Printf("[%s] GAGAL menarik %d. Saldo tidak cukup, Sisa saldo: %d\n", namaUser, jumlah, d.Saldo)
	}
}

func main() {
	fmt.Println("--SIMULASI PERTAHANAN MUTEX--")

	dompet := &DompetDigital2{Saldo: 100000}
	var wg sync.WaitGroup

	// hacker melakukan request bersamaan secara pararel
	wg.Add(2)
	go dompet.TarikUang(80000, "Req_1_Hacker", &wg)
	go dompet.TarikUang(80000, "Req_2_Hacker", &wg)

	wg.Wait()

	// Saldo tetap aman, tidak bisa minus
	fmt.Printf("[AUDIT] Saldo Akhir di Database: %d\n", dompet.Saldo)
}

// RUN MENGGUNAKAN "go run -race bad_toctou.go" untuk mendeteksi race condition
