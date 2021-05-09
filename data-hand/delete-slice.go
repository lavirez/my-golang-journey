package main

import "fmt"

func main() {
	x := []int{12, 31, 6, 9, 102}
	fmt.Println(x)

	// y := []int{234, 77, 88, 99}
	// x = append(x, y...)
	// fmt.Println(x)

	x = append(x[:2], x[4:]...) // this is how you delete from slice
	// here we wanted to delete 2 elements of index 2 and 3
	fmt.Println(x)
}
