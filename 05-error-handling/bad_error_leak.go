package main

import (
	"errors"
	"fmt"
)

// Sentinel Error -> kaku, tigh-coupling
var ErrDBConnection = errors.New("FATAL: Database PostgreSQL di IP 10.0.0.5 port 5432 down")

// Fungsi low level layer (DB)
func CekUserDB(username string) error {
	// simulasi koneksi ke db gagal
	return ErrDBConnection
}

// Fungsi high level layer (API/Controller)
func LoginAPI(username string) {
	fmt.Println("Menerima request dari:", username)
	err := CekUserDB(username)
	if err != nil {
		// tigh-coupling! -> layer api ini harus tahu tentang ErrDBConnection sehingga menghapus batasan antar layer
		if err == ErrDBConnection {
			fmt.Println("Terdeteksi error database!")
		}
		// Information Disclosure karena error langsung di return ke browser user
		fmt.Printf("[RESPON API KE USER] HTTP 300: %v\n", err)
	} else {
		fmt.Println("[RESPON API KE USER] HTTP 200: Login Sukses")
	}
}

func main() {
	fmt.Println("--MENJALANKAN BAD ERROR HANDLING--")
	LoginAPI("hacker_budi")
}
