package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"Jamse Bond", "litterature", "Computer Science"},
		"no_dr":           []string{"being evil", "ice cream", "sunset"},
	}

	x["alavi_alireza"] = []string{"coding", "money", "crypto currency", "Kiana"}

	// now fun thing to know if you user the  _ operator at the range
	// you will get the indexing from the last key in the map run the code
	// like this for i, _ := range x ... you will see
	for i := range x {
		fmt.Printf("%v loves the follwoing:\n", i)
		for j, val := range x[i] {
			fmt.Printf("\t %v %v\n", j, val)
		}
	}
}
