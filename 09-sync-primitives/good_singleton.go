// di sini kita akan menggunakan sync.Once untuk mengatasi bad singleton
// ini adalah Best Practice du Go untuk membuat variabel singleton (variabel yang hanya boleh di buat sekali dalam aplikasi)

package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {
	Status string
}

var (
	dbInstance *Database
	// once time key yg hanya bisa di jalankan sekali
	once sync.Once
)

// Thread-Safe Singleton
func DapatkanKoneksiDBAman() *Database {
	// ini menjamin fungsi di dalamnyna hanya di ekseskusi sekali,
	// berapa kalipun goroutine memanggilnya
	// gorotuine yg datang belakabngan akan otomatis menunggu sampai goroutine pertama
	// selesai membuat goroutine
	once.Do(func() {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("[SUCCESS] Membuka 1 Koneksi TUNGGAL ke Database")
		dbInstance = &Database{Status: "Connected"}
	})
	return dbInstance
}

func main() {
	fmt.Println("--SIMULASI SYNC.ONCE--")
	var wg sync.WaitGroup

	jumlahRequest := 100

	for i := 0; i < jumlahRequest; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DapatkanKoneksiDBAman()
		}()
	}
	wg.Wait()
	fmt.Println("Sistem Selesai. Koneksi stabil, Database aman!")
}
