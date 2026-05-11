package main

type TelemetryData struct {
	DeviceID string
	Status   string
}

// mengembalikan value (salinan), bukan reference (pointer)
// membuat data di function lalu menyalinnya ke main (sharing down)

func ProcessDataValue(id string) TelemetryData {
	data := TelemetryData{
		DeviceID: id,
		Status:   "Active",
	}
	// tidak ada alamat memori yang bocok keluar fungsi
	// variabel "data" yang di buat akan terhapus setelah fungsi di eksekusi dan tak perlu di akses lagi
	// sehingga compiler tak perlu mengirim ke heap dan tetap di stack
	return data
}

func main() {
	// salinan data dari function di kirim ke main dan tersimpan di memori baru di main
	sensorData := ProcessDataValue("Sensor-B")

	_ = sensorData
}

// run menggunakan go run -gcflags="-m" stack_safe.go
// tidak ada output "moved to heap: data"
