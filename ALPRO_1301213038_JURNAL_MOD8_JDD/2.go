package main

import "fmt"

func main() {
	var n, petinju1, petinju2, p1, p2, p3, p4, p5, p6 int
	fmt.Scan(&n)
	petinju1 = 0
	petinju2 = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&p1, &p2, &p3, &p4, &p5, &p6)
		petinju1 += p1 + p2 + p3
		petinju2 += p4 + p5 + p6
	}
	fmt.Println("Petinju A:", petinju1, "dan Petinju B:", petinju2)
	if petinju1 > petinju2 {
		fmt.Println("Pemenang adalah petinju A")
	} else if petinju2 > petinju1 {
		fmt.Println("Pemenang adalah petinju B")
	} else {
		fmt.Println("Draw")
	}
}
