package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	// you cannot do this
	// x = append(x, {123})
	// or this
	// x = append(x, [12, 32, 56])
	fmt.Println(x)
	// now we are going to append the below slice to x 
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...) // use the dots to tell go dump the values in y 
	fmt.Println(x)
}
