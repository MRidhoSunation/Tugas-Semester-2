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

	// field v1 to contain the smallest value of all field values
	for i, min := 0, n[0]; i < len(n); i++ {
		if min > n[i] {
			min = n[i]
		}
		rec.v1 = min
	}

	//field v2 to contain the sum of all the original field values
	rec.vx.v2 = 0
	for _, v := range n {
		rec.vx.v2 += v
	}

	//field v3 to contain the average of the original values of the field
	rec.vx.v3 = rec.vx.v2 / len(n)

	//field v4 to contain the largest value of all the initial values of the field
	for i, max := 0, n[0]; i < len(n); i++ {
		if max < n[i] {
			max = n[i]
		}
		rec.v4 = max
	}
}
