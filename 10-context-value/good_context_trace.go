// ini merupakan penerapan arsitekrut mikroservice dengan membuat tipe data rahasia untuk key context

package main

import (
	"context"
	"fmt"
	"time"
)

// type key custom yang tidak bisa di akses oleh package main
type contextKey string

// konstanta key yang absolut
const traceIDkey contextKey = "trace_id"

// Middleware Layer
func MiddlewareListener(next func(context.Context)) {
	ctx := context.Background()

	// generate unique Trace ID
	traceID := fmt.Sprintf("REQ-%d", time.Now().UnixNano())

	// memasukkan Trace ID ke dalam context with value yg akan mengikuti alur proses data
	ctxDiberiNilai := context.WithValue(ctx, traceIDkey, traceID)

	fmt.Printf("[MIDDLEWARE] Menerima request. Menyisipkan Trace ID: %s\n", traceID)

	// meneruskan ke layer selanjutnya
	next(ctxDiberiNilai)
}

// Controller Layer
func TanganiRequestAPI(ctx context.Context) {
	// controller tak perlu tau soal trace id,
	// hanya meneruskan context.WithValue ke layer selanjutnya
	fmt.Println("[CONTROLLER] Memproses bisnis logic...")
	DatabaseLayer(ctx)
}

// Database Layer
func DatabaseLayer(ctx context.Context) {
	// mengestrak nilai dari context.WithValue
	// wajib melakukan type assertion karena bawaannya adalah any
	traceIDRaw := ctx.Value(traceIDkey)

	var traceID string
	if traceIDRaw != nil {
		traceID = traceIDRaw.(string)
	} else {
		traceID = "UNKNOWN"
	}

	// jika misalnya terjadi error di Database,
	// log ini akan di kirim ke elasticsearch sehingga bisa melacak rute request
	// menggunakan Trace ID
	fmt.Printf("[DATABASE] Menyimpan data. Trace ID: %s\n", traceID)
}

func main() {
	fmt.Println("--SIMULASI TRACEABILITY: GOOD CONTEXT")
	// mensimulasikan satu alur request penuh dari luar ke dalam
	MiddlewareListener(TanganiRequestAPI)
}
