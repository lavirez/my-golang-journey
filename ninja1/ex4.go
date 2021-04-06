package main

import "fmt"

type cus int

var x cus
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 411
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// x = int(x)
	fmt.Println(int(x))
	fmt.Printf("%T\n", int(x))
}