package main

import "fmt"

func main() {
	// array and range
	x := []int{4, 5, 6, 7, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for v := range x {
		fmt.Println(v)
	}

	/*
		this will not work returns: too many variables in range
		for i, j, v := range x {
			fmt.Println(i, j, v)
		}
	*/

}
