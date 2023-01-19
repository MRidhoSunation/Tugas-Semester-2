package main

import "fmt"

func main(){
	var a, b, c, d int
	var p1, p2, c1, c2 int
	fmt.Scanln(&a,&b,&c,&d)
	p1 = permutation(a ,c)
	c1 = combination(a ,c)
	p2 = permutation(b ,d)
	c2 = combination(b ,d)
	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}

func findFactorial(n int, result *int){
	*result = 1
	for i := 1; i <= n; i++ {
		*result = *result * i
	}
}

func permutation(n, r int) int {
	var fn, frn int
	findFactorial(n, &fn)
	findFactorial(n-r, &frn)
	return fn / frn
}

func combination(n, r int)int{
	var fn, frn, fr int
	findFactorial(n, &fn)
	findFactorial(r, &fr)
	findFactorial(n-r, &frn)
	findFactorial(n, &fn)
	return fn / (fr * frn)
}