package main

import "fmt"

var x int // package level scope

func main() {
	// var x int
	// if we used the line 8 code then the scope of x
	// will be the main function
	fmt.Println()
}

func foo() {
	fmt.Println("Hello!")
	x++
	fmt.Println(x)
}
