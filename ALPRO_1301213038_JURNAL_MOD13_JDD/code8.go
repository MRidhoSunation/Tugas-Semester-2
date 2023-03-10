package main

import "fmt"

type arrInt [NMAX]int

func main () {
	var A arrInt
	var i,N,X int

	fmt.Scan(&N,&X)
	for i = 0; i < N; i++{
		fmt.Scan(&A[i])
	}
}

func bin(T arrInt, n, x int)  {
	return search(T,0,n-1,X)
}

func search (T arrInt, left, right, X int) int{
	var mid int = (left + right) / 2
	if left > right {
		return -1
	}
	else if X == T[mid] {
		return mid
	}
	else {
		if x >T[mid] {
			return search(T,left, mid - 1, X)
		}else{
			return search(T,mid+1,right, X)
		}
	}
}