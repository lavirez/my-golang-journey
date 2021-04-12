package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 3; i++{

		fmt.Printf("The outer loop: %d\n", i)
		
		for j:=0; j < 3; j++{
			fmt.Printf("\t \t The outer loop: %d\n", j)
		}
	}
}