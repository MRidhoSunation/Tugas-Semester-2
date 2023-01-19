package main

import "fmt"

func main() {
	var T []int = []int{1, 2, -9, -7, -15, 9999}
	var v = 2
	var n int = len(T)
	Add_data(T, &n)
	n = len(value)
	fmt.Println(FindSekuensial(value, v))
	fmt.Println(Minimum_Value(value, n))
	fmt.Println(Mean_value(value, n))
	fmt.Println(value)
}
/////////////////////////////////////////////////////////////////////////////////

package main

func Meanvalue(tab [9999]int, n int) int {
	var result int
	for _, val := range tab {
		result += val
	}
	return result / n
}