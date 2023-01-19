package main

import "fmt"

type buku struct {
	judul, penulis string
	tahun          int
}

type tabbuku [5]buku

func tambahbuku(kardus *tabbuku, atas *int) {
	if *atas == 4 {
		fmt.Println("kardus penuh")
	} else {
		*atas++
		fmt.Scan(&kardus[*atas].judul, &kardus[*atas].penulis, &kardus[*atas].tahun)
	}
}

func ambilbuku(kardus *tabbuku, atas *int, ambil *buku) {
	if *atas < 0 {
		fmt.Println("kardus kosong")
	} else {
		*ambil = kardus[*atas]
		*atas--
	}
}

func caribuku(kardus *tabbuku, atas *int, x string) {
	var found bool = false
	var ambil buku
	for !found && *atas != -1 {
		ambilbuku(kardus, atas, &ambil)
		fmt.Println("Judul buku yang dikeluarkan:", ambil.judul)
		found = ambil.judul == x
	}
	if found {
		fmt.Println("KETEMU")
	} else {
		fmt.Println("TIDAK KETEMU")
	}
}

func main() {
	var kardus tabbuku
	var atas int

	atas = -1

	tambahbuku(&kardus, &atas)
	tambahbuku(&kardus, &atas)
	tambahbuku(&kardus, &atas)
	tambahbuku(&kardus, &atas)

	caribuku(&kardus, &atas, "C")

	tambahbuku(&kardus, &atas)
	tambahbuku(&kardus, &atas)
	tambahbuku(&kardus, &atas)
	tambahbuku(&kardus, &atas)
	tambahbuku(&kardus, &atas)

	caribuku(&kardus, &atas, "Sistem Digital")
}
