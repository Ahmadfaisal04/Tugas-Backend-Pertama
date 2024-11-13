package main

import (
	"fmt"
)

func main() {
	var operasi int
	var angka1, angka2, hasil int
	fmt.Print("PROGRAM KALKULATOR SEDERHANA DI GOLANG\n1. OPERASI PERKALIAN\n2. OPERASI PEMBAGIAN\n3. OPERASI PENAMBAHAN\n4. OPERASI PENGURANGAN\n5. OPERASI MODULUS\n6. EXIT\nPilih salah satu angka di atas: ")
	fmt.Scan(&operasi)

	if operasi == 1 {
		fmt.Println("OPERASI PERKALIAN")
		fmt.Print("Masukkan angka 1 : ")
		fmt.Scan(&angka1)

		fmt.Print("Masukkan angka 2 : ")
		fmt.Scan(&angka2)
		hasil = angka1 * angka2
		fmt.Println("Hasil : ", hasil)
	} else if operasi == 2 {
		fmt.Println("OPERASI PEMBAGIAN")
		fmt.Print("Masukkan angka 1 : ")
		fmt.Scan(&angka1)

		fmt.Print("Masukkan angka 2 : ")
		fmt.Scan(&angka2)
		hasil = angka1 / angka2
		fmt.Println("Hasil : ", hasil)
	} else if operasi == 3 {
		fmt.Println("OPERASI PENAMBAHAN")
		fmt.Print("Masukkan angka 1 : ")
		fmt.Scan(&angka1)

		fmt.Print("Masukkan angka 2 : ")
		fmt.Scan(&angka2)
		hasil = angka1 + angka2
		fmt.Println("Hasil : ", hasil)
	} else if operasi == 4 {
		fmt.Println("OPERASI PENGURANGAN")
		fmt.Print("Masukkan angka 1 : ")
		fmt.Scan(&angka1)

		fmt.Print("Masukkan angka 2 : ")
		fmt.Scan(&angka2)
		hasil = angka1 - angka2
		fmt.Println("Hasil : ", hasil)
	} else if operasi == 5 {
		fmt.Println("OPERASI MODULUS")
		fmt.Print("Masukkan angka 1 : ")
		fmt.Scan(&angka1)

		fmt.Print("Masukkan angka 2 : ")
		fmt.Scan(&angka2)
		hasil = angka1 % angka2
		fmt.Println("Hasil : ", hasil)
	} else if operasi == 6 {
		fmt.Println("EXIT")
	} else {
		fmt.Println("Input tidak valid, silakan masukkan angka antara 1 hingga 6.")
	}
}