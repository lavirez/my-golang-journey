package main

import "fmt"

//another way rto achieve same thins as in for-state.go file

// for [init]; [cond]; [post] {
//
//}

func main() {
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")
}
