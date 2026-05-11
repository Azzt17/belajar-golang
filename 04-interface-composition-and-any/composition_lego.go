package main

import "fmt"

// ===Interface Segregation Principle===
// membaca saldo
type BalanceReader interface {
	CheckBalance() int
}

// mengubah saldo (menambah dan mengurangi)
type BalanceWriter interface {
	UpdateBalance(amount int)
}

// ===Interface Composition===
// mesin ATM yang butuh fungsi baca dan ubah saldo
type ATM interface {
	BalanceReader // meneruskan fungsi CheckBalance
	BalanceWriter // meneruskan fungsi UpdateBalance
}

// ===Implementasi===
type BankAccount struct {
	Saldo int
}

func (b *BankAccount) CheckBalance() int {
	return b.Saldo
}

func (b *BankAccount) UpdateBalance(amount int) {
	b.Saldo += amount
	fmt.Println("Saldo berhasil diupdate. Saldo sekarang:", b.Saldo)
}

// ===Penggunaan===
// fungsi yang hanya butuh membaca (misalnya app mobile banking ringan), tidak bisa mengupdate saldo
func TampilkanLayar(reader BalanceReader) {
	fmt.Println("[Layar] Saldo anda adalah: Rp", reader.CheckBalance())
}

// fungsi yang butuh baca dan tulis (misalnya mesin ATM)
func TarikTunai(mesin ATM, nominal int) {
	if mesin.CheckBalance() >= nominal {
		fmt.Println("[ATM] Menarik uang sejumlah: Rp", nominal)
		mesin.UpdateBalance(-nominal)
	} else {
		fmt.Println("[ATM] Saldo tidak cukup!")
	}
}

func main() {
	fmt.Println("--SIMULASI INTERFACE COMPOSITION--")

	// Nasabah membuka rekening dengan saldo awal Rp 100.000
	rekeningFarid := &BankAccount{Saldo: 100000}

	// Di Aplikasi Mobile yanf cuma bisa baca:
	TampilkanLayar(rekeningFarid)

	// Di mesin ATM yang bisa baca dan tulis (Composition);
	TarikTunai(rekeningFarid, 20000)

	TampilkanLayar(rekeningFarid)
}
