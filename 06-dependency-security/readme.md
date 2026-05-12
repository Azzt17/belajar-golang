step note:

1. melakukan git mod init di terminal, output: muncul file go.mod kosong
2. menulis kode main.go yang mengakses eksternal library
3. jalankan main.go dengan go run main.go namun error
4. menjalankan go mod tidy di terminal untuk melakukan checksum untuk eksternal library yang di gunakan
5. sekarang file go.mod berisi github.com/fatih/color dan versinya
6. muncul go.sum setelah melakukan go modm tidy yang berisi fingerprint library fatih
7. (Skenario Hacker memanipulasi isi library)mengubah isi file sum bagian fingerprint fatih dari:
    h1:Zp3PiM21/9Ld6FzSKyL5c/BULoe/ONr9KlbYVOfG8+w=
    menjadi:
    h1:Zp3PiM21/9Ld6FzSKedit/BULoe/ONr9KlbYVOfG8+w=
8. menjalankan go mod verify untuk memverifikasi fingerprint di go.sum, output:
      verifying github.com/fatih/color@v1.19.0: checksum mismatch
     downloaded: h1:Zp3PiM21/9Ld6FzSKyL5c/BULoe/ONr9KlbYVOfG8+w=
     go.sum:     h1:Zp3PiM21/9Ld6FzSKedit/BULoe/ONr9KlbYVOfG8+w=

      SECURITY ERROR
      This download does NOT match an earlier download recorded in go.sum.
      The bits may have been replaced on the origin server, or an attacker may
      have intercepted the download attempt.
9. menjalankan program dengan go run main.go, output sama seperti go mod verify
10. mengembalikan isi go.sum seperti semula
11. edit file main.go, tambahkan import: "github.com/google/uuid/"
12. (Simulasi Strict CI/CD) jalankan di terminal: GOFLAGS="-mod=readonly" go build main.go
13. Build akan gagal karena flag ini berfungsi untuk menentukan rule deterministik, jika ada yang tidak sesuai dan perlu di update atau di download (dalam case ini sebuah library baru) maka build akan gagal, ini menuntut developer untuk merapikan dependency terlebih dahulu dengan go mod tidy
