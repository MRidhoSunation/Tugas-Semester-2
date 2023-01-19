package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func main() {
	var A arrInt
	var i, N, total int

	fmt.Scan(&N)
	for i = 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	total = 0
	sum(A, i, &total)
	fmt.Print(total)
}

func sum(T arrInt, i int, total *int) {
	if i >= 0 {
		*total = *total + T[i]
		sum(T, i-1, &*total)
	}
}
