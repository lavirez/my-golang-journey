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

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
