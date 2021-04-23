package main

import "fmt"

func main() {
	fmt.Printf("true && true\t %v \n", true && false)
	fmt.Printf("true && true\t %v \n", true && true)
	fmt.Printf("false || true\t %v \n", false || true)
	fmt.Printf("false || false\t %v \n", false || false)
	fmt.Printf("false || true\t %v \n", false || true)
	fmt.Printf("!false\t \t %v \n", !false)
	fmt.Printf("!true \t \t %v", !true)
}
