package main

import "fmt"

// put whatever returns boolean value
// 2 == 2,  !(2 == 2)

func main() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
}
