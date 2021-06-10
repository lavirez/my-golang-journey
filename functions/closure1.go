package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)
	x++
	fmt.Println(x)
	{
		y := 42
		fmt.Println(y)
	}

	// y++
	// fmt.Println(y)
	// if you uncomment back the above lines then you'd
	// encounter error bacause the scope of y in in the
	// code-block YES you can have code-block in code-block
	// and this is closure
}
