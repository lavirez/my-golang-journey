package main

import "fmt"


func main() {
	x := 4

	y := x << 1

	fmt.Printf("%d\t\t%b\n", x, x)	
	fmt.Printf("%d\t\t%b\n", y, y)
}