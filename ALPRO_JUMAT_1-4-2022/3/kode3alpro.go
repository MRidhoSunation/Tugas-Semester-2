package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, r1 float64
	var x2, y2, r2 float64
	var x, y float64
	var state1, state2 bool
	state1 = true
	state2 = true

	fmt.Scanln(&x1, &y1, &r1)
	fmt.Scanln(&x2, &y2, &r2)
	fmt.Scanln(&x, &y)

	state1 = inside(x1, y1, r1, x, y)
	state2 = inside(x2, y2, r2, x, y)

	if state1 && state2 {
		fmt.Println("Points  inside  circles  1  and  2")
	} else if state1 && !state2 {
		fmt.Println("Points inside circles 1")
	} else if !state1 && state2 {
		fmt.Println("Points inside circles 2")
	} else {
		fmt.Println("Points outside circles 1 and 2")
	}
}

func distance(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func inside(cx, cy, r, x, y float64) bool {
	var d float64
	d = distance(cx, cy, x, y)
	return d <= r

}
