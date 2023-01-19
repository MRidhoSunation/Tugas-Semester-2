package main

import "fmt"

func kabisat(tahun int) bool {
	return (tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0)
}

func mencariTahun(x *int) {
	var cek bool = true
	var tahun int
	fmt.Scanf("%d", &tahun)
	cek = kabisat(tahun)
	for cek == false {
		fmt.Scan(&tahun)
		cek = kabisat(tahun)
	}
	*x = tahun
}

func main() {
	var x int
	mencariTahun(&x)
	fmt.Println(x)
}
