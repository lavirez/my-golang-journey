package main

import "fmt"

func main() {
	x := 42

	if x == 40 {
		fmt.Println("X is 40")
	} else if x == 41 {
		fmt.Println("X was 41")
	} else {
		fmt.Println("X was not 41 or 40")
	}

}
