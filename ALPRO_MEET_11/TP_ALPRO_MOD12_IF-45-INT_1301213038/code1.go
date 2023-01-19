package main

type Trec struct {
	v1 int
	vx struct {
		v2, v3 int
	}
	v4 int
}

func BanyakNilai(rec *Trec) {
	n := []int{rec.v1, rec.vx.v2, rec.vx.v3, rec.v4}

	// field v1 menjadi berisi nilai terkecil dari semua nilai field
	for i, min := 0, n[0]; i < len(n); i++ {
		if min > n[i] {
			min = n[i]
		}
		rec.v1 = min
	}

	//field v2 menjadi berisi jumlah dari semua nilai field semula
	rec.vx.v2 = 0
	for _, v := range n {
		rec.vx.v2 += v
	}

	//field v3 menjadi berisi rata-rata dari nilai-nilai semula dari field tersebut
	rec.vx.v3 = rec.vx.v2 / len(n)

	//field v4 menjadi berisi nilai terbesar dari semua nilai awal field
	for i, max := 0, n[0]; i < len(n); i++ {
		if max < n[i] {
			max = n[i]
		}
		rec.v4 = max
	}
}
