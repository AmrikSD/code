package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	b := Board{
		State: [9]Colour{
			Empty, X, O,
			Empty, X, O,
			Empty, X, O,
		},
	}

	fmt.Println(b)

}
