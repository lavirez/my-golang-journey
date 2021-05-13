package main

import "fmt"

func main() {
	x_1 := []string {"James", "Bond", "Shaken not stirred"}	

	x_2 := []string {"Miss", "Moneypenny", "Heloooo James"}

	mp := [][]string{x_1, x_2}

	fmt.Println(mp)

	for i,v := range mp{
		fmt.Printf("no. %v index of mp slice %v\n", i, v)
		fmt.Println("The indexes are:")	
		for _, val := range mp[i]{
			fmt.Println("\t", val)
		} 
	}
}