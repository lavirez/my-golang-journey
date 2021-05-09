package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 42, 98}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 114, 1013)
	fmt.Println(x)

	// The code bellow will work
	// z := []int{}
	// y := []int{}
	// y = append(z, 89, 12, 30)
	// fmt.Println(y)

	// also

	y := []int{234, 657, 9072, 103}
	x = append(x, y...) // appends the whole y indices :)
	fmt.Println(x)
}
