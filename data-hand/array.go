package main

import "fmt"

func main() {
	var x [5]int // the length of the array is a part of its type
	var y [6]int // this one is totally a different type from x
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

}
