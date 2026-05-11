package main

import "fmt"

func ProcessPaymentSafe(payload map[string]any) {
	fmt.Println("Menerima webhook pembayaran...")

	// Melakukan pengecekan apakah isi any memang adalah int dengan Comma-ok idiom (explicit validation)
	txID, ok := payload["transaction_id"].(int)
	if !ok {
		// jika tipe data tidak sesuai
		fmt.Println("Security Allert: Invalid Payload Type! menolak request ini")
		return
	}
	fmt.Println("Memproses transaksi nomor:", txID)
	fmt.Println("Pembayaran sukses di validasi ke database!")
}

func main() {
	fmt.Println("Menjalankan Webhook...")
	// skenario hacker: mengirim teks(string) ke field angka
	badPayload := map[string]any{
		"transaction_id": "9999; DROP TABLLE users;",
		"status":         "success",
	}
	ProcessPaymentSafe(badPayload)

	fmt.Println("Ini akan di print jika sistem tidak crash dan tetap berjalan")
}
