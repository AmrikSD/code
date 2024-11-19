package main

import (
	"testing"
)

var boardTests = []struct {
	in          Moves
	out         Board
	errExpected bool
	explain     string
}{
	{
		in:      Moves{},
		out:     Board{},
		explain: "An empty array should output a fresh board",
	},
	{
		in: Moves{1, 3},
		out: Board{
			State: [81]Colour{
				Empty, X, Empty, O, Empty, Empty, Empty, Empty, Empty,
				Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
				Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
				Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
				Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
				Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
				Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
				Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
				Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			},
		},
		explain: "a very simple example",
	},
	{
		in:          Moves{-1},
		errExpected: true,
		explain:     "Moves can't be negative",
	},
	{
		in:          Moves{81},
		errExpected: true,
		explain:     "Moves can't be greater than 80",
	},
	{
		in:          Moves{0, 6, 0},
		errExpected: true,
		explain:     "You can't select the same square you've already picked",
	},
}

func TestMoves_ToBoard(t *testing.T) {
	for _, tt := range boardTests {
		t.Run(tt.explain, func(t *testing.T) {
			actual, err := tt.in.ToBoard()

			if tt.errExpected && err == nil {
				t.Fatalf("Expected an error but got none")
			}
			if !tt.errExpected && err != nil {
				t.Fatalf("Got an error when none was expected: %v", err)
			}

			if actual != tt.out {
				t.Errorf("\ngot:\n%v\nwant:\n%v", actual, tt.out)
			}
		})
	}
}
