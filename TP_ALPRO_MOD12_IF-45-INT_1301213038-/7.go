package main

func TerurutB(tab [9999]int, n int) {
	/*
		I.S. defined array of tabs containing n integers.
		F.S. array of sorted tabs shrinks using insertion sort algoritma algorithm
	*/
	var i, j, key int
	for i = 1; i < n; i++ {
		key = tab[i]
		j = i - 1
		for j >= 0 && key > tab[j] {
			tab[j+1] = tab[j]
			j = j - 1
		}
		tab[j+1] = key
	}
}
