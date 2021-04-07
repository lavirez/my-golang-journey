package main

import "fmt"

// print last four years using iota

const (
	a = iota + 2018
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}