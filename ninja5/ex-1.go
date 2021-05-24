package main

import (
	"fmt"
)

type person struct {
	first_name   string
	last_name    string
	fav_icecream []string
}

func main() {
	p1 := person{
		first_name: "Alireza",
		last_name:  "Alavi",
		fav_icecream: []string{
			"Dark Choclate",
			"Vanilla Ice",
		},
	}

	p2 := person{
		first_name: "Marian",
		last_name:  "Alavi Razavi",
		fav_icecream: []string{
			"Strawberry",
			"Saffron",
		},
	}

	for k, v := range p1.fav_icecream {
		fmt.Println(k, v)
	}

	for _, val := range p2.fav_icecream {
		fmt.Println(val)
	}
}
