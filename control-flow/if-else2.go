package main

import "fmt"

// Golang compiler lets you to put multiple shit in the same line
// you just have to user semi colon (;) at the end of the statement

// for instance look at the code below

// note if you save the file the editor will transport the
// fmt.Println("some other statement") to the next line

func main() {
	if x := 42; x < 123 { // we used the ; to both declate
		fmt.Println("001")
	}
	fmt.Println("Here's a statement")
	fmt.Println("some other statement")
}
