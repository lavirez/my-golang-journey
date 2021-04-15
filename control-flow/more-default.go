package main

import "fmt"

func main() {
	switch {
	case (4 == 3):
		fmt.Println("3 is not the same as 4")
	default:
		fmt.Println("this is the defautl case!")
	}
}
