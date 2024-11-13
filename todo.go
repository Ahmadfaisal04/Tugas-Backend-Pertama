package main

import (
	"fmt"
)

func main() {
	var todoList []string
	var pilihan int
	var input string
	var index int

	for {
		fmt.Print(`PROGRAM ToDo - List
				--------------------
				1. Menambah toDo List
				2. Melihat toDo List
				3. Menghapus toDo List
				4. Mengupdate toDo List
				5. Keluar applikasi
				Masukan Angka yang dipilih : `)

		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			fmt.Print("\nMasukkan To-Do: ")
			fmt.Scanln(&input)
			todoList = append(todoList, input)
			fmt.Println("Berhasil menambah todo list!")
		} else if pilihan == 2 {
			if len(todoList) == 0 {
				fmt.Println("Tidak ada todo list yang tersedia.")
			} else {
				fmt.Println("\nDaftar To-Do List: \n")
				for i, todo := range todoList {
					fmt.Printf("%d. %s\n", i+1, todo)
				}
			}
		} else if pilihan == 3 {
			if len(todoList) == 0 {
				fmt.Println("Tidak ada todo list untuk dihapus.")
			} else {
				fmt.Print("\nMasukkan nomor To-Do yang ingin dihapus: ")
				fmt.Scanln(&index)
				if index > 0 && index <= len(todoList) {
					todoList = append(todoList[:index-1], todoList[index:]...)
					fmt.Println("Berhasil menghapus todo list")
				} else {
					fmt.Println("Nomor To-Do tidak valid.")
				}
			}
		} else if pilihan == 4 {
			if len(todoList) == 0 {
				fmt.Println("Tidak ada todo list untuk diupdate.")
			} else {
				fmt.Println("\nDaftar To-Do List:")
				for i, todo := range todoList {
					fmt.Printf("%d. %s\n", i+1, todo)
				}

				fmt.Print("\nMasukkan nomor To-Do yang ingin diupdate: ")
				fmt.Scanln(&index)
				if index > 0 && index <= len(todoList) {
					fmt.Print("Masukkan To-Do baru: ")
					fmt.Scanln(&input)
					todoList[index-1] = input
					fmt.Println("Berhasil mengupdate todo list")
				} else {
					fmt.Println("Nomor To-Do tidak valid.")
				}
			}
		} else if pilihan == 5 {
			fmt.Println("Terima kasih! Keluar dari aplikasi.")
			return
		} else {
			fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		}
	}
}
