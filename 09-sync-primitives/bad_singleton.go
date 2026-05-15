// ini adalah skenario banyak koneksi database dalam 1 waktu
// jika koneksi di inisialisasi di setiap awal proses,
// maka akan ada 1000 permintaan koneksi ke database ketika ada 1000 proses bersamaan
// ini bisa menynebabkan Connection Pool Exhaustion, keadaan ketika database menolak koneksi
// atau crash karena kelebihan beban
//

package main

import (
	"fmt"
	"sync"
	"time"
)

// simualsi object database
type Database struct {
	Status string
}

var dbInstance *Database

// kesalahan arsitektur:
// memungkinan banyak koneksi dibuat bersamaan
func DapatkanKoneksiDB() *Database {
	if dbInstance == nil {
		// simulasi waktu koneksi (10 milidetik)
		// saat goroutine sleep di sini, goroutine lain akan masuk dan mengeksekusi ini juga
		time.Sleep(10 * time.Millisecond)

		fmt.Println("[WARNING] Membuka koneksi TCP/IP Baru ke Database...")
		dbInstance = &Database{Status: "Connected"}
	}
	return dbInstance
}

func main() {
	fmt.Println("--SIMULASI CONNEECTION EXHAUSTION--")
	var wg sync.WaitGroup

	jumlahRequest := 100

	for i := 0; i < jumlahRequest; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DapatkanKoneksiDB() // ini akan di ulangi sebanyak jumlah request
		}()
	}

	wg.Wait()
	fmt.Println("Sistem Selesai. Database mungkin sudah Crash!")
}
