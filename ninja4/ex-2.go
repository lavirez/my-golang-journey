package main

import "fmt"

func main() {
	// so the difference between array and slice
	// is that in the first one the has specific
	// length but the latter has nothing
	x := []int{42, 34, 56, 76, 32}

	for _, v := range x {
		fmt.Println(v)
	}
}
