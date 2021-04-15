package main

import "fmt"

// when you use value after switch it no longer switches on the bool
// it swithes on the value

func main() {
	switch "James" {
	case (4 == 2):
		fmt.Println("This is a bool value so bye bye")
	case "Jamse", "Bond", "Eine":
		fmt.Println("i can see the value James here")
	case "black":
		fmt.Println("I can see no value of Jamse here")
	}
}
