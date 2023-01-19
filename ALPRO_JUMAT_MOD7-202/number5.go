package main

import "fmt"

func main() {
	var nama, nama_Min, nama_Max string
	var p int
	var faktor, jumlah, min, max, rata float64
	nama_Min = ""
	nama_Max = ""
	fmt.Scan(&nama)
	for nama != "STOP" {
		fmt.Scan(&p)
		jumlah = 0
		for i := 0; i < 3*p; i++ {
			fmt.Scan(&faktor)
			jumlah += faktor
		}
		rata = jumlah / float64(3*p)
		if nama_Min == "" || nama_Min != "" && min > rata {
			min = rata
			nama_Min = nama
		}
		if nama_Max == "" || nama_Max != "" && max < rata {
			max = rata
			nama_Max = nama
		}
		fmt.Scan(&nama)
	}
	fmt.Printf("%s %.2f\n", nama_Max, max)
	fmt.Printf("%s %.2f\n", nama_Min, min)
}
