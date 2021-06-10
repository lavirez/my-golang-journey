package main

import "fmt"

// anynomouse struct

func main() {
	tr1 := struct {
		door      int
		color     string
		fourWheel bool
	}{
		door:      3,
		color:     "Orange",
		fourWheel: true,
	}

	fmt.Println(tr1)
}
