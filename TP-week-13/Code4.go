package main

import "fmt"

func Convert_desimalTobiner(number int) {
	if number == 1 {
		fmt.Print("1")
		return
	} else {
		Convert_desimalTobiner(number / 2)
		fmt.Print(number % 2)
	}
}

func main() {
	Convert_desimalTobiner(13)
}
