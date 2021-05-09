package main

import "fmt"

func main() {
	m := map[string]int{
		"Jamse Bond":      34,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m["Jamse Bond"])

	// so now we add the element
	m["todd"] = 33

	fmt.Println(m)
}
