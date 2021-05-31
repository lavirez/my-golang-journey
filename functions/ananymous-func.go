package main

import "fmt"

func main() {
	func(){
		fmt.Println("Ananymous fuction has run")
	}()

	func(x int){
		fmt.Println("the meaning of life", x)
	}(42)
}