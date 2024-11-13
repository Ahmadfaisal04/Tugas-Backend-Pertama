package main

import (
	"fmt"
)

func main() {
	var todoList []string
	var pilihan int

	for {
		showMenu()
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			todoList = addTodo(todoList)
		case 2:
			viewTodos(todoList)
		case 3:
			todoList = deleteTodo(todoList)
		case 4:
			todoList = updateTodo(todoList)
		case 5:
			fmt.Println("Terima kasih! Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		}
	}
}

// showMenu menampilkan menu utama aplikasi
func showMenu() {
	fmt.Print(`PROGRAM ToDo - List
--------------------
1. Menambah toDo List
2. Melihat toDo List
3. Menghapus toDo List
4. Mengupdate toDo List
5. Keluar aplikasi
Masukan Angka yang dipilih : `)
}

// addTodo untuk menambah to-do ke dalam daftar
func addTodo(todoList []string) []string {
	var input string
	fmt.Print("\nMasukkan To-Do: ")
	fmt.Scanln(&input)
	todoList = append(todoList, input)
	fmt.Println("Berhasil menambah todo list!")
	return todoList
}

// viewTodos untuk melihat semua todo dalam daftar
func viewTodos(todoList []string) {
	if len(todoList) == 0 {
		fmt.Println("Tidak ada todo list yang tersedia.")
	} else {
		fmt.Println("\nDaftar To-Do List: \n")
		for i, todo := range todoList {
			fmt.Printf("%d. %s\n", i+1, todo)
		}
	}
}

// deleteTodo untuk menghapus todo dari daftar
func deleteTodo(todoList []string) []string {
	if len(todoList) == 0 {
		fmt.Println("Tidak ada todo list untuk dihapus.")
		return todoList
	}

	var index int
	fmt.Print("\nMasukkan nomor To-Do yang ingin dihapus: ")
	fmt.Scanln(&index)

	if index > 0 && index <= len(todoList) {
		todoList = append(todoList[:index-1], todoList[index:]...)
		fmt.Println("Berhasil menghapus todo list")
	} else {
		fmt.Println("Nomor To-Do tidak valid.")
	}
	return todoList
}

// updateTodo untuk mengupdate todo dalam daftar
func updateTodo(todoList []string) []string {
	if len(todoList) == 0 {
		fmt.Println("Tidak ada todo list untuk diupdate.")
		return todoList
	}

	viewTodos(todoList)

	var index int
	fmt.Print("\nMasukkan nomor To-Do yang ingin diupdate: ")
	fmt.Scanln(&index)

	if index > 0 && index <= len(todoList) {
		var input string
		fmt.Print("Masukkan To-Do baru: ")
		fmt.Scanln(&input)
		todoList[index-1] = input
		fmt.Println("Berhasil mengupdate todo list")
	} else {
		fmt.Println("Nomor To-Do tidak valid.")
	}
	return todoList
}
