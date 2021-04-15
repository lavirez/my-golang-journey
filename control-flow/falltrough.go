package main

import "fmt"

func main() {
	switch {
	case (4 == 3):
		fmt.Println("3 is not the same as 4")
	case (3 == 2):
		fmt.Println("3 is also not equal to 2")
	case (2 == 2):
		fmt.Println("So 2 is equal to 2 ")
		fallthrough // you put this the next case will run no matter what
	case (4 == 1):
		fmt.Println("But what the heck 4 is not 1! ")
	case (4 == 4):
		fmt.Println("WHAT THE FUCK? where is 4 equals to 4?!")

	}
}
