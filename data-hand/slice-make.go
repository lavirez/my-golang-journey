package main

import "fmt"

func main() {
	x := make([]int, 9, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 12
	x[1] = 43

	x = append(x, 9)
	fmt.Println(x)

	// if you increase the len of the x more
	// than the make given capacity the cap
	// will be doubled to the given cap
	x = append(x, 21)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
