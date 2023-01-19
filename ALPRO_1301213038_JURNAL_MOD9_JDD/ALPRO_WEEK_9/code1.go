package main

import (
	"fmt"
	"math"
)

type points struct {
	x, y float64
}

func main() {
	var P1, P2 points
	var dis float64
	fmt.Scanln(&P1.x, &P1.y, &P2.x, &P2.y)
	dis = distance(P1, P2)
	fmt.Println("The distance between point (", P1.x, ",", P1.y, ") and (", P2.x, ",", P2.y, ") is ", dis)
}

func distance(P1, P2 points) float64 {
	var dis float64
	dis = (P2.x-P1.x)*(P2.x-P1.x) + (P2.y-P1.y)*(P2.y-P1.y)
	dis = math.Sqrt(dis)
	return dis
}
