package main

import "fmt"

func main() {
	for x := 0; x < 100; x++ {
		if x%2 == 0 {
			if x%5 == 0 {
				continue
			} else if x%3 == 0 {
				continue
			} else {
				fmt.Println(x)
			}
		}
	}
}
