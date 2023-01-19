package main

import "fmt"

type arr [5]int

func printArray(ar arr, n int) int {
	var size int = n
	extensive := len(ar)
	if size == extensive {
		return 0
	}
	fmt.Printf("%d ", ar[size])
	return printArray(ar, size+1)

}

func main() {
	var arr arr
	arr[0] = 21
	arr[1] = 22
	arr[2] = 23
	arr[3] = 24
	arr[4] = 25
	printArray(arr, 0)
}
