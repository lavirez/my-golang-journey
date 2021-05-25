package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person 
	ltk bool 
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Jamse",
			last: "Bond",
			ltk: true,
		}
	}
}