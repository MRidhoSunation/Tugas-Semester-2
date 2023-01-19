package main

import "fmt"

func ManyDigit(number int, total int) int {
	var temp int = total
	if number > -1 && number < 10 {
		temp++
		return temp
	}
	var num int = number
	num = (num - (num % 10)) / 10
	temp++
	return ManyDigit(num, temp)
}

func main() {
	var number int = 12345
	result := ManyDigit(number, 0)
	fmt.Println(result)
}
