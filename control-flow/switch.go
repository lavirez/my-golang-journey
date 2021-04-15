package main

import "fmt"

func main() {
	x := 42
	switch {
	case x == 2:
		fmt.Println("Javan Inc. is a fraud!")
	case x == 32:
		fmt.Println("Those Mother Fuckers in the Javan Inc. are Motherfucking frauds")
	case x == 42:
		fmt.Println("Holu MOther of GOD")
		// you could use falltrough here so that the next case will be
		// printed out
	case x == 42:
		fmt.Println("This one wont get printed")
	}
}
