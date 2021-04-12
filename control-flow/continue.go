package main

import "fmt"

// print the even numbers before 100

func main() {
	x := 0
	for {
		x++

		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}
		fmt.Println(x)

	}

	fmt.Println("Boom./")
}
