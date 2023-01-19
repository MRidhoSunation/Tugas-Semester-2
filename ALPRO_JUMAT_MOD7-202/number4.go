package main

import "fmt"

func main() {
	var a, total, before, dec int
	var rata float64
	dec = 0
	total = 0
	a = 0
	i := 0
	for dec != 3 {
		i++
		before = a
		fmt.Print("Hari ", i, " : ")
		fmt.Scanln(&a)
		for a < 0 || a > 100 {
			fmt.Print("Hari ", i, " : ")
			fmt.Scanln(&a)
		}
		if before > a {
			dec += 1
		}
		total += a
	}
	rata = float64(total) / float64(i)
	fmt.Println("Museum buka selama", i, "hari")
	fmt.Printf("Rata-rata pengunjung %.2f\n", rata)
}
