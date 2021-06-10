package main

import "fmt"

// struct is a compoite data type aka aggregate data type

type person struct{
	first string
	last string
	age int
}

func main(){
	p1 := person{
		first: "James",
		last: "Bond",
		age: 32,
	}

	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		age: 26,
	}

	// so we created something like class (type person) and then some
	// instances of it ... well no this is a good aproximation but we
	// never use those words in golang instead we user we created values
	// of type person that's it 

	fmt.Printf("%T\t %v \n",p1, p1)
	fmt.Printf("%T\t %v \n", p2, p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}