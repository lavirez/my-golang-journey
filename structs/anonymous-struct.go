package main

import "fmt"

// the ananymous struct is a struct that doesn't decalre name
// so that you would user it once so that the memory will be free

func main(){
	p1 := struct{
		first string
		last string 
		age int
	}{
		first: "Jamse",
		last: "Bond",
		age: 32,
	}

	fmt.Println(p1)
}