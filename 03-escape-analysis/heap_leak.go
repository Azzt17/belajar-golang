package main

type TelemetryData struct {
	DeviceID string
	Status   string
}

// case: sharing up (return pointer)
// fungsi ini membuat data di dalam func dan mereturn alamatnya ke main

func ProcessDataPointer(id string) *TelemetryData {
	data := TelemetryData{
		DeviceID: id,
		Status:   "Active",
	}
	// karena datanya di simpan di variabel local dan akan di akses lagi setelah fungsi selesai
	// compiler terpaksa memindahkannya ke heap
	return &data
}

func main() {
	// misal ini di call 10.000 kali per detik
	sensorData := ProcessDataPointer("Sensor-A")
	// di kosongkan agar tidak memicu escape tambahan dari fungsi Println
	_ = sensorData
}

// run menggunakan go run -gcflags="-m" heap_leak.go
// untuk melihat escape-analysis nya (tool bawaan go)
