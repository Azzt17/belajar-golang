// simulasi banyak pesanan dalam waktu yg sama, eg: flash sale event
// setiap pesanan membuat goroutine baru (bad concurrency)
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func ProsesPesanan(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// simulasi query ke db
	time.Sleep(1 * time.Second)
}

func main() {
	fmt.Println("--Memulai Flash Sale--")
	var wg sync.WaitGroup
	jumlahPesananMasuk := 5000 // 5000 pesanan dalam 1 detik
	fmt.Printf("Goroutine awal: %d\n", runtime.NumGoroutine())

	// menciptakan goroutine sebanyak jumlah request
	for i := 1; i <= jumlahPesananMasuk; i++ {
		wg.Add(1)
		go ProsesPesanan(i, &wg)
	}
	fmt.Printf("Gorotuine saat flash sale berlangsung: %d\n", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Flash Sale Selesai.")
}
