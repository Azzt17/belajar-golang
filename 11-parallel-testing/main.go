// ini adalah program yang akan di test menggunakan mekanisme testing bawaan go
package main

import (
	"fmt"
	"time"
)

// Sistem Bank Sederhana yang TIDAK memiliki Mutex
type DompetDigital struct {
	Saldo int
}

// Transfer memindahkan uang dari dompet ke tujuan.
// Mengembalikan error (string) jika saldo tidak cukup.
func (d *DompetDigital) Transfer(jumlah int, tujuan string) string {
	// TIME-OF-CHECK
	if d.Saldo >= jumlah {
		// Simulasi jeda pemrosesan database/jaringan
		time.Sleep(5 * time.Millisecond)

		// TIME-OF-USE
		d.Saldo -= jumlah
		return "Sukses"
	}
	return "Saldo tidak cukup"
}

func main() {
	dompet := &DompetDigital{Saldo: 100}
	hasil := dompet.Transfer(50, "User_B")
	fmt.Printf("Status: %s, Sisa Saldo: %d\n", hasil, dompet.Saldo)
}
