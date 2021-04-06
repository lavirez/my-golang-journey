package main

import "fmt"

// so here we goin to use iota (IOTA) to accomplish the same thing that the output  
// of bit-shift2.go file is 	

const (
	// the following line takes care of the first iota which is
	_ = iota
	//kb = 1024
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}