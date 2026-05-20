# 🧪 Golang Technical Lab & Architecture Playground

**Author:** Farid Wajdi

Repositori ini bukan sekadar kumpulan tutorial dasar, melainkan sebuah laboratorium teknis yang mendokumentasikan perjalanan eksplorasi mendalam terhadap arsitektur internal Go (Golang). 

Fokus utama dari *playground* ini adalah membedah bagaimana Go menangani memori, pola konkurensi skala besar, mitigasi kerentanan tingkat sistem (seperti *Race Conditions*), dan praktik observabilitas (*profiling* & *testing*).

## 🏗️ Struktur Laboratorium

Repositori ini dibagi ke dalam modul-modul terisolasi yang menguji konsep spesifik:

### Fase 1: Fondasi & Manajemen Memori
* `01` s.d `07` - **Core Mechanics:** Eksplorasi fundamental Go, *Escape Analysis* (Heap vs Stack), implementasi *Interface*, dan deteksi *Goroutine Leaks*.

### Fase 2: Sistem Konkurensi & Sinkronisasi
* `08-concurrency-patterns` - **Worker Pools:** Membangun sistem pemrosesan paralel (seperti simulasi *Flash Sale*) menggunakan arsitektur *Communicating Sequential Processes* (CSP) via *Channels*.
* `09-sync-primitives` - **Memory Synchronization:** Mitigasi kerentanan *Time-of-Check to Time-of-Use* (TOCTOU) menggunakan `sync.Mutex`, `sync.Map`, dan `atomic` operations.
* `10-context-value` - **Request Lifecycle:** Membedah `context.Context` sebagai pembawa sinyal pembatalan lintas-goroutine (*timeout/deadline*) dan koper diplomatik (*Context Value*) yang aman.

### Fase 3: Observabilitas, Testing & DevSecOps
* `11-parallel-testing` - **Race Condition Hunting:** Memanfaatkan Table-Driven Tests dan `t.Parallel()` dipadukan dengan flag `-race` untuk mendeteksi tabrakan memori secara otomatis.
* `12-testability-mocking` - **Dependency Injection:** Menerapkan *Repository Pattern* dan *Mocking* (`testify`) untuk menciptakan kode yang 100% terisolasi dan *testable* tanpa memerlukan koneksi database fisik.
* `13-profiling-benchmark` - **X-Ray Profiling:** Penggunaan `go test -bench` dan `go tool pprof` untuk mendeteksi *Memory Leaks* dan menganalisis beban *Garbage Collector* (mencegah Application-Layer DoS).

## 🚀 Cara Menggunakan Repositori Ini

Setiap folder dirancang sebagai modul yang berdiri sendiri. Untuk menjalankan eksperimen atau pengujian di modul tertentu:

1. Masuk ke direktori yang diinginkan:
    ```bash
    cd 13-profiling-benchmark
2. Jalankan simulasi atau tes dengan flag analitik:
    ```bash
    go test -v -race       # Untuk audit keamanan memori
    go test -bench=.       # Untuk audit performa dan alokasi RAM

## 🎯 Tujuan Pembuatan

Eksperimen di dalam repositori ini dirancang untuk memastikan bahwa sistem perangkat lunak tidak hanya "berfungsi", tetapi juga aman dari desain (Secure-by-Design), efisien dalam penggunaan resource, dan deterministik di bawah tekanan eksekusi paralel.
