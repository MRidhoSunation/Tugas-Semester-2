package main
import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	konversi(a, &b)
	fmt.Println(b)
}

func pangkat(a, b int) int {
	var hasil int = 1
	for i := 1; i <= b; i++ {
		hasil *= a
	}
	return hasil
}

func konversi(biner int, desimal *int) {
	var tampung int
	*desimal = 0
	i := 0
	for biner > 0 {
		tampung = biner % 10
		*desimal += tampung * pangkat(2, i)
		biner = biner / 10
		i++
	}
}