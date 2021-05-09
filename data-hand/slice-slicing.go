package main

import "fmt"

func main() {
	x := []int{4, 5, 68, 42, 90}
	fmt.Println(x[1:])
	fmt.Println(x[2:4])

	for i, v := range x {
		fmt.Println(i, v)
	}

	// here is how you can do the the above
	// without the range clause

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
