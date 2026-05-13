// case kode ini adalah ada user yang memanggil API ini dan requestnya di teruskan ke db yg lambat
// namun sebelum selesai user membatalkan requestnya (Timeout)

package main

import (
	"fmt"
	"runtime"
	"time"
)

// fungsi dlm API
func PanggilApi(ch chan string) {
	// simulasi db yg lambat
	time.Sleep(3 * time.Second)

	// setelah 3 detik, fungsi ini mencoba mengirim data ke ch
	// namun sudah tdk ada yg menerima sehingga goroutine ini akan selamanyna berjalan di ini
	ch <- "Data User"

	fmt.Println("Goroutine selesai dengan normal (ini tak akan pernah tercetak)")
}

// endpoint API
func EndpointHandler() {
	ch := make(chan string)

	// membuat goroutine baru
	go PanggilApi(ch)

	// mekanisme timeout (agar user tidak menunggu terlalu lama)
	select {
	case res := <-ch:
		fmt.Println("Berhasil mendapat data:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("[API RESPONS] HTTP 408: Request Timeout. User membatalkan.")
		// fungsi ini selesai dan return
		// namun lupa memberitahu goroutine utk berhenti
	}
}

func main() {
	fmt.Println("--Simulasi DOS (goroutine leak)--")
	fmt.Println("Jumlah go rotuine awal (hanya fungsi main:", runtime.NumGoroutine())

	// user melakukan request
	EndpointHandler()

	// tunggu 4 detik untuk melihat apakah goroutine di API akan mati sendiri
	fmt.Println("Menunggu 4 detik memantau background process")
	time.Sleep(4 * time.Second)

	// bukti memori leak (angkanya bertambah):
	fmt.Println("Jumlah Goroutine akhir:", runtime.NumGoroutine())
	fmt.Println("jika endpoint ini di serang 1  juta request, server akan kehabisan ram dan down")
}
