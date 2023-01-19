package main

import "fmt"

const max int = 30

type IntArray struct {
	tabInt [max]int
	n      int
}

func main() {
	var array1, array2 IntArray
	fmt.Print("input array pertama: ")
	inputArray(&array1)
	fmt.Print("input array kedua: ")
	inputArray(&array2)
	sortArray(&array1)
	sortArray(&array2)
	// cetakData(array1)
	// cetakData(array2)
	fmt.Print(isSimilar(array1, array2))
}

func inputArray(T *IntArray) {
	var arr int
	i := 0
	arr = 1
	for i < max && arr != 0 {
		fmt.Scan(&arr)
		T.tabInt[i] = arr
		i++
		T.n++
	}
}

func sortArray(T *IntArray) {
	var pass, idx, i, temp int
	pass = 1
	for pass <= T.n-1 {
		idx = pass - 1
		i = pass
		for i < T.n {
			if T.tabInt[idx] > T.tabInt[i] {
				idx = i
			}
			i++
		}
		temp = T.tabInt[pass-1]
		T.tabInt[pass-1] = T.tabInt[idx]
		T.tabInt[idx] = temp
		pass++
	}
}

// func cetakData(T IntArray) {
// 	for i := 0; i < T.n; i++ {
// 		fmt.Print(T.tabInt[i], "\n")
// 	}
// }

func isSimilar(T1, T2 IntArray) bool {
	var cek bool
	cek = true
	if T1.n == T2.n {
		for i := 0; i < T1.n && cek; i++ {
			if T1.tabInt[i] != T2.tabInt[i] {
				cek = false
			}
		}
	} else {
		cek = false
	}
	return cek
}
