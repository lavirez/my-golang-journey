package main

import (
	"fmt"
)

const (
	a = 42 //untyped constant
	v int = 43 //typed constant
)

func main() {
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	
	fmt.Println(a)
	fmt.Printf("%T", a)
}