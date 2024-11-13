package main

import (
	"fmt"
)

type akun struct {
	password string
	saldo    float64
	history  []string
}

func main() {
	var akun = akun{
		password: "112233",
		saldo:    50000,
	}

	var inputPassword string
	var pilihan int

	for {
		fmt.Print(`
--------------------
WELCOME TO CLI BANK
--------------------
1. Masuk Aplikasi
2. Keluar Aplikasi
Masukkan angka 1/2: `)
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&inputPassword)

			if inputPassword == akun.password {
				fmt.Println(`
--------------------
Login Berhasil.....
--------------------`)
				mainMenu(&akun)
			} else {
				fmt.Println("Password salah!")
			}

		} else if pilihan == 2 {
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			break
		} else {
			fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		}
	}
}

func mainMenu(akun *akun) {
	var pilihan int

	for {
		fmt.Print(`
1 Lihat Saldo
2 Top Up Saldo
3 Kirim Uang
4 Tarik Uang
5 Lihat History Transaksi
6 Exit
Masukkan Pilihan : `)
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			fmt.Printf("Saldo Anda saat ini: Rp%.2f\n", akun.saldo)
		} else if pilihan == 2 {
			var amount float64
			fmt.Print("Masukkan jumlah top up: Rp")
			fmt.Scanln(&amount)
			if amount > 0 {
				akun.saldo += amount
				transaksi := fmt.Sprintf("Top up: Rp%.2f", amount)
				akun.history = append(akun.history, transaksi)
				fmt.Printf("Top up sebesar Rp%.2f berhasil!\n", amount)
			} else {
				fmt.Println("Jumlah top up harus lebih besar dari 0.")
			}
		} else if pilihan == 3 {
			var penerima string
			var amount float64
			fmt.Print("Masukkan username penerima: ")
			fmt.Scanln(&penerima)
			fmt.Print("Masukkan jumlah uang yang ingin dikirim: Rp")
			fmt.Scanln(&amount)
			if amount <= akun.saldo && amount > 0 {
				akun.saldo -= amount
				transaksi := fmt.Sprintf("Kirim uang ke %s: Rp%.2f", penerima, amount)
				akun.history = append(akun.history, transaksi)
				fmt.Printf("Kirim uang sebesar Rp%.2f ke %s berhasil!\n", amount, penerima)
			} else if amount <= 0 {
				fmt.Println("Jumlah kirim uang harus lebih besar dari 0.")
			} else {
				fmt.Println("Saldo Anda tidak cukup untuk transaksi ini.")
			}
		} else if pilihan == 4 {
			var amount float64
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp")
			fmt.Scanln(&amount)
			if amount <= akun.saldo && amount > 0 {
				akun.saldo -= amount
				transaksi := fmt.Sprintf("Tarik uang: Rp%.2f", amount)
				akun.history = append(akun.history, transaksi)
				fmt.Printf("Tarik uang sebesar Rp%.2f berhasil!\n", amount)
			} else if amount <= 0 {
				fmt.Println("Jumlah tarik uang harus lebih besar dari 0.")
			} else {
				fmt.Println("Saldo Anda tidak cukup untuk transaksi ini.")
			}
		} else if pilihan == 5 {
			fmt.Println("\nHistory Transaksi:")
			if len(akun.history) == 0 {
				fmt.Println("Tidak ada transaksi.")
			} else {
				for i, transaksi := range akun.history {
					fmt.Printf("%d. %s\n", i+1, transaksi)
				}
			}
		} else if pilihan == 6 {
			fmt.Println("Keluar dari aplikasi...")
			return
		} else {
			fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		}
	}
}
