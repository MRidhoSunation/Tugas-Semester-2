package main

import "fmt"

func main(){
	var a,b,c int
	fmt.Scanln(&a,&b,&c)
	var fogoh, gohof, hofog int
	fogoh = fx(gx(hx(a)))
	gohof = gx(hx(fx(b)))
	hofog = hx(fx(gx(c)))
	fmt.Println(fx(gx(hx(a))))
	fmt.Println(gx(hx(fx(b))))
	fmt.Println(hx(fx(gx(c))))
	fmt.Println(fogoh, gohof, hofog)
}

func fx(x int)int{
	return x*x
}

func hx(x int)int{
	return x + 1
}

func gx(x int) int{
	return x - 2
}
