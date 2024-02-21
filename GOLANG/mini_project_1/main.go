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

func main() {
	id := 0
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
			for _, book := range books {
				fmt.Println(book.id)
				fmt.Println("Nama:", book.nama)
				fmt.Println("Deskripsi:", book.desc)
				fmt.Println("Status", book.status)
				fmt.Println(" ")
			}
		case "c":
			s.Scan()
			deleteid := s.String()
			for i, book := range books {
				if book.id == deleteid {
					books = append(books[:i])
				}
			}
		}
	}

}
