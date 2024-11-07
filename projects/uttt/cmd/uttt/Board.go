package main

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
	State [9]Colour
}

func (b Board) String() string {
	var result string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result += "[" + b.State[i*3+j].String() + "]"
			if j < 2 {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}
