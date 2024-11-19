package main

import (
	"errors"
	"fmt"
)

type Colour int

func (c Colour) String() string {
	switch c {
	case Empty:
		return " "
	case X:
		return "X"
	case O:
		return "O"
	default:
		return "Unknown"
	}
}

const (
	Empty Colour = iota
	X
	O
)

type Board struct {
	State  [81]Colour
	Winner Colour
}

func (b Board) String() string {
	var result string
	switch b.Winner {
	case Empty:
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				result += "[" + b.State[i*9+j].String() + "]"
				if j < 2 {
					result += " "
				}
			}
			result += "\n"
		}
		return result

	case X:
		result = "X WON LOL"
		return result
	case O:
		result = "O WON LOL"
		return result
	default:
		panic("yeah this shouldn't happen.")
	}
}

type Moves []int

func (moves Moves) ToBoard() (Board, error) {
	b := Board{}
	turn := X
	for idx, move := range moves {
		if move < 0 || move > 80 {
			return b, errors.New("moves must be between '0' and '80'")
		}
		if b.State[move] != Empty {
			return b, errors.New(fmt.Sprintf("on turn %d, %s tried to play into tile %d but it had already been played", idx+1, turn, move))
		}

		b.State[move] = turn
		if turn == X {
			turn = O
		} else {
			turn = X
		}
	}
	return b, nil
}
