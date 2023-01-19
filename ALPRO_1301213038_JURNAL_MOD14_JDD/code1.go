package main

import "fmt"

var (
	pita                    string
	CC, CC2                 byte
	EOP, EOP2               bool
	indeks, number, number1 int
	a, b, total             bool
	test                    string
)

func main() {
	fmt.Scanln(&pita)
	START()
	if !EOP {
		test = spilter()
		a = test == "DORA" || test == "SE" || test == "IT"
	}
	for !EOP {
		test = spilter()
		test = test + "."
		number = 0
		number1 = 0
		START2()
		for !EOP2 {
			number1++
			if DIGIT2() {
				number++
			}
			ADV2()
		}
		b = number == number1 && number > 0
		spilter()
	}
	total = a && b
	fmt.Println(total)
}

func START() {
	indeks = 0
	CC = pita[0]
	EOP = CC == '.'
}
func START2() {
	indeks = 0
	CC2 = test[indeks]
	EOP2 = CC2 == '.'
}

func ADV() {
	indeks = indeks + 1
	CC = pita[indeks]
	EOP = CC == '.'
}

func ADV2() {
	indeks = indeks + 1
	CC2 = test[indeks]
	EOP2 = CC2 == '.'
}

func DIGIT() bool {
	return CC >= 48 && CC <= 57
}
func DIGIT2() bool {
	return CC2 >= 48 && CC2 <= 57
}

func spilter() string {
	var part string
	for CC == 45 {
		part = part + string(CC)
		ADV()
	}
	part = string(CC)
	if DIGIT() {
		ADV()
		for DIGIT() {
			part = part + string(CC)
			ADV()
		}
	} else if CC >= 65 && CC <= 90 || CC >= 97 && CC <= 122 {
		ADV()
		for DIGIT() || CC >= 65 && CC <= 90 || CC >= 97 && CC <= 122 {
			part = part + string(CC)
			ADV()
		}
	}
	return part
}
