package main

import (
	"errors"
	"fmt"
)

// Internal sentinel error
var ErrTimeout = errors.New("koneksi timeout")

// Low layer (DB)
func AmbilDataUser() error {
	// Error di database, di wrapped dengan konteks
	return fmt.Errorf("gagal query ke tabel users: %w", ErrTimeout)
}

// Middle layer (Service/Business logic)
func ValidasiLogin() error {
	err := AmbilDataUser()
	if err != nil {
		// tidak mengecek err == ErrTimeout, tidak mempedulikan apa isi errornya, langsung wrap ke atas
		return fmt.Errorf("proses validasi terhenti: %w", err)
	}
	return nil
}

// High layer (API/Controller)
func LoginEndpoint() {
	fmt.Println("Menerima request login...")

	err := ValidasiLogin()
	if err != nil {
		// pisahkan log internal dan response eksternal yang di kirim ke users
		fmt.Printf("[LOG SERVER INTERNAL] Alert: %v\n", err)

		// errors.Is untuk unwrapped error agar konteksnya jelas
		if errors.Is(err, ErrTimeout) {
			// Opaque response ke users untuk mencegaj data leak
			fmt.Println("[RESPONSE API KE USER] HTTP 503: Layanan sedang sibuk, silahkan coba lagi nanti.")
		} else {
			// Memberikan respons global
			fmt.Println("[RESPONES API KE USER] HTTP 500: Terjadi kesalahan pada server.")
		}
	}
}

func main() {
	fmt.Println("--MENJALANKAN GOOD ERROR HANDLING--")
	LoginEndpoint()
}
