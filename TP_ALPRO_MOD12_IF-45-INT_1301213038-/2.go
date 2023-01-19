package main

import "fmt"

func TambahData(tab *[9999]int, n *int) {
	/*
				    I.S. input data is ready on the input device
					F.S. tab array contains a number of n integers from input, input ends when number
		is 9999
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
