package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	tr1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Orange",
		},
		fourWheel: true,
	}

	sd1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Grey",
		},
		luxury: false,
	}

	fmt.Println(sd1)
	fmt.Println(tr1)
	fmt.Println(sd1.doors)
}
