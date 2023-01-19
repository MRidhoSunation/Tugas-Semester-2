package main

import "fmt"

var pita string
var CC byte
var EOP bool
var indeks int

func main() {
	var valid bool
	valid = false

	fmt.Scan(&pita)
	START()
	if CC == 'D' {
		ADV()
		if CC == 'O' {
			ADV()
			if CC == 'R' {
				ADV()
			}
			if CC == 'A' {
				ADV()
			}
			if CC == '-' {
				ADV()
				valid = EOP
			}
		}
	}
}

func START() {
	indeks = 0
	CC = pita[indeks]
	EOP = CC == '.'
}
func ADV() {
	indeks++
	CC = pita[indeks]
	EOP = CC == '.'
}

func DIGIT() bool {
	return CC >= '0' && CC <= '9'
}
