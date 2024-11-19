package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	b := Board{
		State: [81]Colour{
			Empty, X, O,
			Empty, X, O,
			Empty, X, O,
		},
	}

	moves := Moves{
		1, 2, 4, 5, 7, 8, 8,
	}

	fmt.Println(moves.ToBoard())
	fmt.Println(b)

}
