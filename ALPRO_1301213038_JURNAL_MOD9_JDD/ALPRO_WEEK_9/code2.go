package main

import "fmt"

type Gudang [256]int

func main() {
	var Gudang1 Gudang = [256]int{1, 2, 3, 4, 5}
	reverseArr(&Gudang1)
	fmt.Println(Gudang1)
	cekPalindrom := checkPalindrome(Gudang1)
	fmt.Println(cekPalindrom)

}

func reverseArr(arr *Gudang) {
	var temp Gudang
	for j := 0; j < 256; j++ {
		temp[j] = arr[j]
	}
	for i := 255; i >= 0; i-- {
		arr[i] = temp[255-i]
	}
}

func checkPalindrome(arr Gudang) bool {
	for i := 0; i <= 255; i++ {
		if arr[i] != arr[255-i] {
			return false
		}
	}
	return true
}

func fillArr(arr *Gudang, fill string) {
	if fill == "Odd" {
		for i := 0; i%2 == 0 && i <= 255; i++ {
			arr[i] = i + 1
		}
	} else if fill == "Even" {
		for i := 0; i%2 != 0 && i <= 255; i++ {
			arr[i] = i + 1
		}
	}
}
