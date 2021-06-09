package main

import "fmt"

func main() {
	fmt.Println(foo())
	bar()
}

func foo() string {
	s := "Hello World"
	return s
}

func bar() func() int {
	return func() int {
		return 42
	}
}
