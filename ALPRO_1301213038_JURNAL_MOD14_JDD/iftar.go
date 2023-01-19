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
		if CC == 'B' {
			ADV()
			if CC == '-' {
				ADV()
				for DIGIT() {
					ADV()
				}
				if CC == '-' {
					ADV()
					if DIGIT() {
						for DIGIT() {
							ADV()
						}
						valid = EOP
					}
				}
			}
		}
	}
	fmt.Print(valid)
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
