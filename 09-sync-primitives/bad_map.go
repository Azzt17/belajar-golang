// di go, map bawaan (map[string]string) tidak kebal terhadap konkurensi, jika ada dua
// goroutine yang mencoba write/read secara bersamaan, akan menyebabkan Fatal Error: concurrent map writes/read
// + ini tdk bisa menggunakan recover()
// satu saja map di aplikasi yg tidak di lock bisa menciptakaan celah DoS yg fatal

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("--SIMULASI SERVER CRASH: CONCURRENT MAP--")

	// map standar bawaan go
	sessionCache := make(map[string]string)
	var wg sync.WaitGroup

	jumlahLoginBersamaan := 100

	// 100 user mencoba login bersamaan
	for i := 0; i < jumlahLoginBersamaan; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			token := fmt.Sprintf("token_rahasia_%d", id)
			userKey := fmt.Sprintf("user_%d", id)

			// write ke dalam map yg sama tanppa lock
			sessionCache[userKey] = token
		}(i)
	}

	wg.Wait()
	fmt.Println("Server berhasil bertahan? ini hanya akan ter print jika server aman")
}
