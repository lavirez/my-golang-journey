package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%#x\n", a)
	b := a << 1
	fmt.Printf("%d\n%b\n%#x\n", b, b, b)
}