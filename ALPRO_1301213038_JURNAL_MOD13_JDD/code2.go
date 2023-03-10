package main

import "fmt"

var count int = 0

func CountDigits(num int) int {
	if num > 0 {
		count++
		CountDigits(num / 10)
	}
	return count
}

func main() {
	var num int = 0
	var result int = 0

	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)

	result = CountDigits(num)
	fmt.Printf("Count of digits is: %d\n", result)
}
