package main

import "fmt"

const NMAX = 1000

type arrint [NMAX]int

func main() {
	var A arrint
	var i, N int

	fmt.Scan(&N)
	for i = 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	print(A, N)
}

func print(T arrint, n int) {
	// master function (optional)
}

func printArray(T arrint, i, n int) {
	// recursive function
	if i < n {
		fmt.Printf("%v", T[i])
		printArray(T, i+1, n)
	}
}
