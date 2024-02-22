package main

import (
	"bufio"
	"fmt"
	"os"
)

type buku struct {
	id     int
	nama   string
	desc   string
	status string
}

func printBooks(books []buku) {
	for _, book := range books {
		fmt.Println(" ")
		fmt.Println("no:", book.id)
		fmt.Println("Nama:", book.nama)
		fmt.Println("Deskripsi:", book.desc)
		fmt.Println("Status", book.status)
	}
}

func main() {
	id := 1
	var books []buku
	var first string
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("a = tambah buku, b = liat data buku-buku, c = hapus buku")
		s.Scan()
		first = s.Text()
		switch first {
		case "a":
			var nama, desc, status string
			status = "Tersedia"
			fmt.Println("nama buku:")
			s.Scan()
			nama = s.Text()
			fmt.Println("deskripsi buku:")
			s.Scan()
			desc = s.Text()
			books = append(books, buku{
				id:     id,
				nama:   nama,
				desc:   desc,
				status: status,
			})
			id++
		case "b":
			printBooks(books)
		case "c":
			printBooks(books)
			fmt.Println("Masukkan nomor buku yang ingin dihapus...")
			s.Scan()
			deleteid := s.Text()
			for i, book := range books {
				if fmt.Sprint(book.id) == deleteid {
					books = append(books[:i], books[:i+1]...)
				}
			}
		case "d":
			printBooks(books)
			fmt.Println("Masukkan nomor buku yang ingin diedit...")
			s.Scan()
			editid := s.Text()
			for i, book := range books {
				if fmt.Sprint(book.id) == editid {
					books = append(books[:i], books[:i]...)
				}
			}
		}
	}

}
