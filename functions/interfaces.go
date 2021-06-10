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

type human interface{
	speak()
}

func bar (h human){
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

// func (r reciever) identifier(parameters) (return(s)) { code... } 
 
func (s secretAgent) speak() {
	fmt.Println("I am, ", s.first, s.last, "from a secretAgent view")
}

func (p person) speak() {
	fmt.Println("I am, ", p.first, p.last, "-- From a noraml person view")
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Jamse",
			last: "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last: "Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		"Dr.",
		"Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
}