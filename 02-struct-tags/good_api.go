package main

import (
	"encoding/json"
	"fmt"
)

// menggunakan struct tags untuk pembatasan keamanan
type SecureUserProfile struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"` // menginstruksikan encoder utk mengabaikan field func init() {
	WalletSaldo  int    `json:"-"`
}

func main() {
	fmt.Println("Menjalankan API..")
	secureUser := SecureUserProfile{
		ID:           102,
		Username:     "farid_w",
		PasswordHash: "bcrypt127475417...rahasia",
		WalletSaldo:  5000000,
	}

	fmt.Println("Sistem mengkonversi data ke JSON untuk di kirim ke aplikasi")

	jsonData, err := json.MarshalIndent(secureUser, "", " ")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	// data rahasia terenkapsulasi
	fmt.Println("Respons API yang di kirim ke internet:")
	fmt.Println(string(jsonData))
}
