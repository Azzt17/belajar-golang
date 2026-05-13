// di sini kita akan menggunakan "context" untuk kill switch, jika endpoint mati atau batal,
// maka semua go routine di dalamnya akan serentak di shutdown untuk membersihkan RAM

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// fungsi API
func PanggilApi(ctx context.Context, ch chan string) {
	select {
	case <-time.After(3 * time.Second):
		// skenario API selesai normal (sangat lambat)
		ch <- "Data User"
	case <-ctx.Done():
		// jika endpoint di batalkan, ini akan aktif (exit path)
		fmt.Println("[LOG INTERNAL] Sinyal batal diterima. Goroutine di shutdown untuk membersihkan RAM")
		return
	}
}

// konsumen
func EndpointHandler() {
	// membuat context dengan batas 1 detik
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // untuk membersihkan sisa memori contect ini sedniri

	ch := make(chan string)

	go PanggilApi(ctx, ch) // membuat go routine dan menitipkan context kedalamnya

	select {
	case res := <-ch:
		fmt.Println("Berhasil mendapat data:", res)
	case <-ctx.Done():
		// menangkap sinyal timeout dari context
		fmt.Println("[API RESPONS] HTTP 408: Request Timeout. User membatalkan request")
	}
}

func main() {
	fmt.Println("--Simulasi Pertahanan Memori menggunakan Context--")
	fmt.Println("Jumlah goroutine awal:", runtime.NumGoroutine())

	EndpointHandler()

	fmt.Println("Menunggu 4 detik memantau backgorund process...")
	time.Sleep(4 * time.Second)

	// bukti memori aman (goroutine kembali ke angka awal)
	fmt.Println("Jumlah goroutine akhir: ", runtime.NumGoroutine())
}
