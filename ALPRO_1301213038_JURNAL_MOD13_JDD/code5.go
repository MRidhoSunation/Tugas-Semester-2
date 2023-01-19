package main

import "fmt"

func main() {
	var x, y, z, total int

	total = 1
	fmt.Scanln(&x, &y)
	z = y
	pow(x, y, z, &total)
	fmt.Println(total)
}

func pow(x, y, z int, total *int) {
	if y > 0 {
		*total = *total * x
		pow(x, y-1, z, &*total)
	}
	if z == 0 {
		*total = 1
	}
}
