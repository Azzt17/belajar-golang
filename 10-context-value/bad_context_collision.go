// ini mensimulasikan app yang menggunakan middleware buatan sendiri,
// namun menggunakan library logger eksternal yang kebetulan menggunakan
// key string yg sama :"id" sehingga menyebabkan collision

package main

import (
	"context"
	"fmt"
)

// simulasi library eksternal
func LoggerMiddleware(ctx context.Context) context.Context {
	// library ini menyisipkan ID log ke dlm context menggunakan key "id"
	return context.WithValue(ctx, "id", "LOG-999")
}

// Database Layer
func simpanKeDatabase(ctx context.Context) {
	userID := ctx.Value("id")

	fmt.Printf("[DATABASE] Menyimpan data untuk User ID: %v\n", userID)
}

func main() {
	fmt.Println("--SIMULASI KEY COLLISTION--")

	ctx := context.Background()

	// 1. Middleware auth memasukkan id user dengan tipe key string
	ctx = context.WithValue(ctx, "id", "USER-123")
	fmt.Println("[AUTH] User berhasil login. ID di sisipkan ke context")

	// 2. Request melewati middleware eksternal (logger)
	ctx = LoggerMiddleware(ctx)

	// 3. Request sampai ke Database
	simpanKeDatabase(ctx)
}
