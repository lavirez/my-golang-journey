package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("This is my first func expression!")
	}
	f()

	g := func(x int) {
		fmt.Println("The year big brother started watching!", x)
	}
	g(1984)
}
