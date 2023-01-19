package main

// function to swap both values
func tukar(a *int, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

func TerurutA(tab *[9999]int, n int) {
	/*
		I.S. defined array of tabs containing n integers.
		F.S. the sorted tab array is enlarged using the selection sort algorithm
	*/
	var i, j, key int
	for i = 0; i < n; i++ {
		key = i
		for j = i + 1; j < n; j++ {
			if tab[j] < tab[key] {
				key = j
			}
		}
		if tab[i] != tab[key] {
			tukar(&tab[i], &tab[key])
		}
	}
}
