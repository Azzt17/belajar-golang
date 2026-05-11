package main

import "fmt"

// fungsi yang menerima apa saja:
func ProcessPayment(payload map[string]any) {
	fmt.Println("Menerima webhook pembayaran...")
	// programmer yakin kalau transaction_id pasti adalah integer
	// sehingga langsung melakukan type assertion tanpa validasi (blind type assertion)
	txID := payload["transaction_id"].(int)

	fmt.Println("Memproses transaksi nomor:", txID)
	fmt.Println("Pembayaran sukses di validasi ke database!")
}

func main() {
	fmt.Println("--Menjalankan Webhook--")
	// skenario normal: user mengirim angka (int)
	PayloadA := map[string]any{
		"transaction_id": 1055,
		"status":         "success",
	}
	ProcessPayment(PayloadA)

	fmt.Println("Menerima webhook pembayaran...")
	// skenario hacker: mengirim teks(string) SQL Injection ke field angka
	PayloadB := map[string]any{
		"transaction_id": "9999; DROP TABLE users;",
		"status":         "success",
	}
	ProcessPayment(PayloadB)

	// cek status sistem
	fmt.Println("Baris ini akan di print ketika sistem masih aktif")
}
