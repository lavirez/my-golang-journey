package main

import "fmt"

func main(){
	states_slice := make([]string, 50, 50)
	states_slice = []string{
		"Alaska",
		"Alabama",
		"Arkansas",
		"Arizona",
		"California",
		"Colorado",
		"Connecticut",
		"Delaware",
		"Florida",
		"Georgia",
		"Hawaii",
		"Iowa",
		"Idaho",
		"Illinois",
		"Indiana",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Massachusetts",
		"Maryland",
		"Maine",
		"Michigan",
		"Minnesota",
		"Missouri",
		"Mississippi",
		"Montana",
		"North Carolina",
		"North Dakota",
		"Nebraska",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"Nevada",
		"New York",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"Pennsylvania",
		"Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Virginia",
		"Vermont",
		"Washington",
		"Wisconsin",
		"West Virginia",
		"Wyoming",
	}

	fmt.Println(states_slice)
	fmt.Println(cap(states_slice))
	fmt.Println(len(states_slice))

	for i := 1; i < len(states_slice); i++{
		fmt.Println(i+1, states_slice[i])
	}
}