package main

import "fmt"

func main() {
	fmt.Println(4*)
	a := factorial(4)
	fmt.Println(a)
}


func factorial(n int) int {
	return n * factorial(n-1)
}