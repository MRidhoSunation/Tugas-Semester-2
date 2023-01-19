package main

import "fmt"

func TambahData(tab *[9999]int, n *int) {
	/*
		    I.S. data masukan telah siap pada piranti masukan
			F.S. array tab berisi sejumlah n bilangan bulat dari masukan,masukan berakhir apabila bilangan
			adalah 9999
	*/
	var data int
	for {
		fmt.Scan(&data)
		if data != 9999 {
			tab[*n] = data
			*n++
		} else {
			break
		}
	}
}
