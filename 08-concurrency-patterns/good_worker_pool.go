// ini menerapkan pattern FAN-IN FAN-OUT untuk best case arsitektur dlm mengatasi potensi DoS
// karena goroutine leak maupun memory bloat
// membatasi jumlah goroutine, menggunakan buffered channel untuk antrian
// untuk menerapkan "Dont Communicate by Sharing Memory, Sharing Memory by Communicate"

package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// goroutine worker
func Kasir(ctx context.Context, id int, antrean <-chan int, hasil chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// multiplexer (listening antrean dan sinyal shutdown untuk graceful shutdown)
		select {
		case pesanan, ok := <-antrean:
			if !ok {
				// jika channel di close, ok bernilai false, maka worker tahu bahwa sedang tidak ada request
				// kemudian masuk dlm keadaan sleep untuk hemat resource memory
				fmt.Printf("[LOG] worker %d di nonaktifkan, antrian di tutup.\n", id)
				return
			}

			// memproses pesanan (simulasi db 1 detik)
			time.Sleep(1 * time.Second)
			hasil <- fmt.Sprintf("Pesanan #%d di selesaikan oleh worker %d", pesanan, id)

		case <-ctx.Done():
			// sinyal untuk maintanance (grateful shutdown)
			fmt.Printf("[ALERT] Worker %d di nonaktifkan (server shutdown).\n", id)
			return
		}
	}
}

func main() {
	fmt.Println("--Memulai Flash Sale--")

	const jumlahPesanan = 5000
	const jumlahKasir = 100 // pembatasan jumlah goroutine aktif dlm satu waktu

	// membuat CSP
	antrean := make(chan int, jumlahPesanan)  // buffered channel utk antrean masuk
	hasil := make(chan string, jumlahPesanan) // buffered channel utk antrean keluar

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Printf("Goroutine awal: %d\n", runtime.NumGoroutine())

	// FAN OUT Pattern: menggunakan goroutine dlm jumlah tetap terus menerus
	for i := 1; i <= jumlahKasir; i++ {
		wg.Add(1)
		go Kasir(ctx, i, antrean, hasil, &wg)
	}

	// memasukkan pesanan ke dalam channel
	for i := 1; i <= jumlahPesanan; i++ {
		antrean <- i
	}

	// menutup input channel: tidak ada request baru yang boleh masuk lagi ke dlm antrian
	// ini akan membuat worker masuk dlm keadaan sleep
	close(antrean)

	// check jumlah goroutine
	fmt.Printf("Gorotuine saat Flash Sale Berlangsung: %d\n", runtime.NumGoroutine())

	// FAN-IN Pattern: mengumpulkan hasil pemrorsesan oleh goroutine dlm antrean
	go func() {
		wg.Wait()    // menunggu semua worker selesai
		close(hasil) // tutup channel
	}()

	// Print hasil
	count := 0
	for res := range hasil {
		if count < 5 {
			fmt.Println("->", res) // print 5 pesanan pertama
		} else if count%500 == 0 {
			fmt.Printf("[LOG] ... %d pesanan telah di proses\n", count)
		}
		count++
	}

	fmt.Println("Flash Sale Selesai, Semua pelanggan terlayanai tanpa memicu DoS.")
}
