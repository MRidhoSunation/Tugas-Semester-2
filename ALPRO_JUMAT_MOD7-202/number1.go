package main

import "fmt"

func main() {
	var positif, negatif, banyak, x, y int
	fmt.Scan(&banyak)
	for i := 0; i < banyak; i++ {
		fmt.Scan(&x, &y)
		if x > 0 {
			positif += x
		} else {
			negatif += x
		}
		if y > 0 {
			positif += y
		} else {
			negatif += y
		}
	}
	fmt.Println("Negative:", negatif)
	fmt.Println("Positive:", positif)
}
