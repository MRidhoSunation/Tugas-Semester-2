package main

import "fmt"

func main() {
	var m, n, menit int
	fmt.Scanln(&m, &n)
	menit = 0
	for {
		menit++
		if menit%m == 0 && menit%n == 0 {
			break
		}
	}
	fmt.Println(menit)
}
