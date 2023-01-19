package main

import "fmt"

const nmax int = 2022

type student struct {
	name, sid string
	gpa       float64
}

type tabMhs [nmax - 1]student

func main() {
	var tabel tabMhs
	var n int
	fmt.Scanln(&n)
	inputData(&tabel, n)
	urutAscending(&tabel, n)
	fmt.Print("----------------------------\n")
	fmt.Print("Pengurutan secara ascending:\n")
	cetakData(tabel, n)
}

func inputData(t *tabMhs, n int) {
	var name, sid string
	var gpa float64
	i := 0
	for i < n {
		fmt.Scan(&name, &sid, &gpa)
		t[i].name = name
		t[i].sid = sid
		t[i].gpa = gpa
		i++
	}

	/*
		for i := 0; i < n; i++ {
			fmt.Scanln(&t[i].name, &t[i].sid, &t[i].gpa)
		)
		versi lain */
}

func cetakData(t tabMhs, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %.2f \n", t[i].name, t[i].sid, t[i].gpa)
	}
}

func urutAscending(t *tabMhs, n int) {
	var pass, i, idx int
	var temp student
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if t[idx].gpa > t[i].gpa {
				idx = i
			}
			i++
		}
		//tukar
		temp = t[pass-1]
		t[pass-1] = t[idx]
		t[idx] = temp
		//iterasi
		pass++
	}
}
