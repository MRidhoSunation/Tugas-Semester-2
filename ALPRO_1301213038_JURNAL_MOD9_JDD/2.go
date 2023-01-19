package main

import "fmt"

type tabgoal [2022]int

func inputData(t *tabgoal, n *int) {
	var menang int
	fmt.Scan(&menang)
	*n = 0
	for menang >= 0 {
		t[*n] = menang
		*n++
		fmt.Scan(&menang)
	}
}

func rataan(t tabgoal, n int) float64 {
	var jum int

	for i := 0; i < n; i++ {
		jum += t[i]
	}
	return float64(jum) / float64(n)
}

func main() {
	var t tabgoal
	var n int

	for i := 0; i < 3; i++ {
		inputData(&t, &n)
		fmt.Println(rataan(t, n))
	}
}
