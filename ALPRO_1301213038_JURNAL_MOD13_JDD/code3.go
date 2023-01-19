package main

import "fmt"

func binary(ang int) {
	if ang > 1 {
		binary(ang / 2)
	}
	fmt.Print(ang % 2)
}

func main() {
	var ang int
	fmt.Scanln(&ang)
	binary(ang)
}
