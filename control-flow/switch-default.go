package main

import "fmt"

// so here we got no true case in the switch scope that's why the
// default case will be runned
// also falltrough has no effect int he below blocks

func main() {
	switch {
	case "James Bond is Fag":
		fmt.Println("We all know that's not true James!")
	case false:
		fmt.Println("Pffft the false case abosolutely not gonna run! like ever")
	case (4 == 3):
		fmt.Println("3 is not the same as 4")
	case (3 == 2):
		fmt.Println("3 is also not equal to 2")
	case (2 == 1):
		fmt.Println("So 2 is not equal to 1 ")
		fallthrough // this falltrough has no effect
	case (4 == 1):
		fmt.Println("But what the heck 4 is not 1! ")
	case (4 == 2):
		fmt.Println("4 is not the same as 2 of coues!")
	default:
		fmt.Println("this is the defautl case!")
	}
}
