package main

import "fmt"

func main() {
	m := map[string]int{
		"James Bond":      34,
		"Miss Moneypenny": 27,
		"Todd":            33,
	}

	for k, v := range m { // key and value
		fmt.Println(k, v)
	}

	xi := []int{34, 23, 98, 56, 23}

	for i, v := range xi { // index and value
		fmt.Println(i, v)
	}
}
