package main

import "fmt"

func main() {
	// we are going through multi dimensional arrays
	jb := []string{"James", "Bond", "Choclate", "martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
