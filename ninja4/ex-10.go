package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"Jamse Bond", "litterature", "Computer Science"},
		"no_dr":           []string{"being evil", "ice cream", "sunset"},
	}

	delete(m, "no_dr")
	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v)
	}

}
