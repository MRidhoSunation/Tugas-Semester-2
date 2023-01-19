package main

func NilaiMinimum(tab [9999]int, n int) int {
	// Returns the location of the minimum value contained in an array of tabs containing n integers
	var min int
	for i := 0; i < n; i++ {
		if tab[i] < tab[min] {
			min = i
		}
	}
	return min
}
