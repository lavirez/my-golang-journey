package main

import "fmt"

// the foo function will be executed after bar meaning after program closes
 
func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Im foo")
}


func bar() {
	fmt.Println("Im bar")
}
