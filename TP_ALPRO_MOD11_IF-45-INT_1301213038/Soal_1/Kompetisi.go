package main

import (
	"fmt"
)

type set [1000000]int

var arr set

func main() {
	var angka, temp_mid int
	var hasil float64
	var i int = 0
	fmt.Scan(&angka)
	for angka != -5313541 && angka >= 0 {
		if angka == 0 {
			sortArr(&arr, i)
			panjang := i
			if panjang%2 == 0 {
				temp_mid = (panjang / 2)
				hasil = (float64(arr[temp_mid]) + float64(arr[temp_mid-1])) / 2
				fmt.Println(hasil)
			} else {
				temp_mid = (panjang / 2)
				hasil = float64(arr[temp_mid])
				fmt.Println(hasil)
			}
			fmt.Scan(&angka)
		} else {
			arr[i] = angka
			fmt.Scan(&angka)
			i++
		}
	}
}

func sortArr(ar *set, cek int) {
	for i := 0; i < cek-1; i++ {
		for j := 0; j < cek-i-1; j++ {
			if ar[j] < ar[j+1] {
				temp := ar[j]
				ar[j] = ar[j+1]
				ar[j+1] = temp
			}
		}
	}

}
