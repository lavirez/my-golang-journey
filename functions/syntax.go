package main

import "fmt"

func main() {
	foo()
	bar("James")
	s := woo("Moneypenny")
	fmt.Println(s)
	x, y := mouse("Ian", "Felmming")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receiver) identifier(parameters) (returns(s)) { ... }

func foo() {
	fmt.Println("hello Im foo")
}

// know the difference with parameter and arguments
// we define our func with parameters (if any)
// we call our func and pass in arguments (if any)

func bar(s string) {
	fmt.Println("hello", s)
}

// everythin in Go is PASS BY VALUE

//
func woo(s string) string {
	return fmt.Sprint("hello from woo,", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, say "hello"`)
	b := false
	return a, b
}
