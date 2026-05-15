// ada 2 solusi utama untuk mengatasi konkurensi pada map:
// 1. Standard Mutex Map (di gunakan untuk 90% case harian)
// 2. sync.Map (digunakan hanya jika map tersebut bertindak sebagai Cache yang
// jarang di update, tetapi di baca ribuan kali per detik oleh banyak goruitine yg berebeda)

package main

import (
	"fmt"
	"sync"
)

// SOLUSI 1: Map biasa + RWMurex
type StandardSessionCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *StandardSessionCache) Set(key, value string) {
	c.mu.Lock() // mekanisme lock agar hanya 1 proses write yg boleh terjadi
	c.data[key] = value
	c.mu.Unlock()
}

func (c *StandardSessionCache) Get(key string) (string, bool) {
	c.mu.RLock() // mengizinkan Read secara bersamaan
	defer c.mu.RUnlock()
	val, exists := c.data[key]
	return val, exists
}

// SOLUSI 2: sync.Map (khusus cache Read-Heavy)
// sync.Map tidak butuh struct tambahan/inisialisasi 'make'
var specializedCache sync.Map

func main() {
	fmt.Println("--SIMULASI SAFE MAP--")
	var wg sync.WaitGroup
	jumlahLogin := 100

	// inisialisasi solusi 1
	standarCache := &StandardSessionCache{
		data: make(map[string]string),
	}

	for i := 0; i < jumlahLogin; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			key := fmt.Sprintf("user_%d", id)
			val := fmt.Sprintf("token_%d", id)

			// 1. menggunakan standar Mutex
			standarCache.Set(key, val)

			// 2. menggunakan sync.map
			specializedCache.Store(key, val)
		}(i)
	}
	wg.Wait()

	fmt.Println("[AUDIT] Sistem berhasil bertahan mengatasi konkurensi")

	tokenStandar, _ := standarCache.Get("user_50")
	fmt.Println("Token dari Standar Mutex:", tokenStandar)

	// X-ray sync.Map, karena menggunakan any, maka wajib melakukan type assertion
	tokenSyncMapRaw, _ := specializedCache.Load("user_50")
	tokenSyncMap := tokenSyncMapRaw.(string) // type assertion

	fmt.Println("Token dari sync.Map:", tokenSyncMap)
}
