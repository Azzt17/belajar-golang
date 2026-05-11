package main

import (
	"encoding/json"
	"fmt"
)

// 1 struct untuk semua (ambil data dari db, respon request API,dsb)
type UserProfile struct {
	ID           int
	Username     string
	PasswordHash string // important data
	WalletSaldo  int    // important data
}

func main() {
	fmt.Println("Menjalankan API...")
	userFromDB := UserProfile{
		ID:           101,
		Username:     "farid_w",
		PasswordHash: "bcrypt$2a$afejanfkae..rahasia",
		WalletSaldo:  500000,
	}
	fmt.Println("Mengkonversi data JSON untuk di kirim ke apliaksi mobile...")

	// Marshal adalah encoder yg mengubah struct memori ke teks json
	jsonData, err := json.MarshalIndent(userFromDB, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Kalau respons API nya di cegat bisa nge leak PasswordHash dan WalletSaldo
	fmt.Println("Respons API yang dikirim ke internet:")
	fmt.Println(string(jsonData))
}
