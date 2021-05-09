package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]

	fmt.Println(v)
	fmt.Println(ok)
	// you can use v in every if statement

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("So the barnabas is not in the map", v)
	} else if v, ok := m["James"]; ok {
		fmt.Println("so it means that 'James' is in the map", v)
	}

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("and aslo the miss monneypenny is in the map too", v)
	}

}
