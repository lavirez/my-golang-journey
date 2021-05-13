package main

import "fmt"

func main(){
	x := map[string][]string{
		"bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"Jamse Bond", "litterature", "Computer Science"},
		"no_dr": []string{"being evil", "ice cream", "sunset"},
	}
	fmt.Println(x)
}