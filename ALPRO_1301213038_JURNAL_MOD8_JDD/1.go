package main

import "fmt"

func main() {
	var d1, d2, ganjil int
	var stop bool
	fmt.Scanln(&d1, &d2)
	ganjil = 1
	stop = d1%2 == 0 && d2%2 == 0
	for !stop {
		if d1%2 != 0 && d2%2 == 0 {
			ganjil++
		}
		fmt.Scanln(&d1, &d2)
		stop = d1%2 == 0 && d2%2 == 0
	}
	fmt.Println(ganjil)
}
