package main 

import "fmt"

const (
	 a = iota
	 b = iota
	 c = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", b)

}

// relust will be the a = 0 , b = 1, c = 2
// and all the type will be outputted as int