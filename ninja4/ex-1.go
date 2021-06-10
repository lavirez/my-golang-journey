package main

import "fmt"

func main() {
	ar := [5]int{1, 2, 5, 42, 8}

	for _, v := range ar {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", ar)
}
