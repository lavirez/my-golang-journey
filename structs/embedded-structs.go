package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

// fact: this ==> secreAgent{}  or ==> person{} both are composite literal

type secertAgent struct{
	person
	first string // this one is to demonstrate the name collision
	ltk bool // just so you know ltk is acronym for license to kill :D
}

func main(){
	sa1 := secertAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 32,
		},
		first: "something to demonstrate collision",
	}

	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		age: 26,
	}

	fmt.Printf("%T\t %v \n",sa1, sa1)
	fmt.Printf("%T\t %v \n", p2, p2)

	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	// no the collision
	fmt.Println(sa1.person.first, "\t this is of type person")
	fmt.Println(sa1.first, "\t this is of type secreAgent and not related to person")
}