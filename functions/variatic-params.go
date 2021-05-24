package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5, 7, 8, 9)
	fmt.Println("The total is ", x)
}

/// func (r reciever) identifier(parameter(s)) (return(s)) { code }

func sum(x ...int) int {
	fmt.Println(x)        // slice of int
	fmt.Printf("%T\n", x) // type slice of int

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index postion ", i, " we are now adding", v, " to the total ", sum)

	}
	return sum
}
