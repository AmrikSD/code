package main

import (
    "fmt"
    "math/rand"
)

type BoardMap [9]bool
func (bm BoardMap) isTickTacToe() bool {
    var cells = [8][3]int{
        {0,1,2},
        {3,4,5},
        {6,7,8},

        {0,3,6},
        {1,4,7},
        {2,5,8},

        {0,4,8},
        {6,4,2},
    }
    for _, cellList := range cells {
        c := 0
        for _, cell := range cellList {
            if bm[cell] {
                c++
            }
        }
        if c == 3 {
            return true
        }
    }
    return false;
}

type Board struct {
	Checkmate bool
    CanPlay   bool
	XMap      BoardMap
	OMap      BoardMap
}
func NewBoard() Board {
    return Board{
        CanPlay: true,
    }
}
func (b Board) isFull() bool {
    c := 0
    for i := 0; i < len(b.XMap); i++ {
        if b.XMap[i] || b.OMap[i] {
            c++
        }
    }
    return c == len(b.XMap)
}

func NewBoards() [9]Board {
 	var boards [9]Board
	for i := range boards {
		boards[i] = NewBoard()
	}
	return boards
}

type Game struct {
	Turn   int // turn mod 2 == 0 -> player 1...
	Boards [9]Board
	XMap   BoardMap
	OMap   BoardMap
}
func NewGame() Game {
    return Game{
        Boards: NewBoards(),
    }
}

func (game *Game) print() {
	fmt.Printf("Turn:\t%d\n", game.Turn)
	fmt.Printf("Meta board:")
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Println()
		}
		if game.XMap[i] {
			fmt.Print("[X]")
		} else if game.OMap[i] {
			fmt.Print("[O]")
		} else {
			fmt.Print("[ ]")
		}
	}
    fmt.Printf("\n\nOpen Boards:")
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Println()
		}
        if game.Boards[i].CanPlay {
            fmt.Print("[▨]")
        }else {
            fmt.Print("[ ]")
        }
	}

    fmt.Printf("\n\nBig board:\n")

	for rowBlock := 0; rowBlock < 3; rowBlock++ {
		for row := 0; row < 3; row++ { 
			for colBlock := 0; colBlock < 3; colBlock++ {
				for col := 0; col < 3; col++ {
					boardIndex := rowBlock*3 + colBlock
					cellIndex := row*3 + col
					board := game.Boards[boardIndex]
					if board.OMap[cellIndex] {
						fmt.Print("[O]")
					} else if board.XMap[cellIndex] {
						fmt.Print("[X]")
					} else {
						fmt.Print("[ ]")
					}
				}
				fmt.Print(" ")
			}
			fmt.Println() 
		}
		fmt.Println() 
	}
}

// Bool (winning move), Error (why you can't play somewhere)
func (g *Game) TakeTurn(pickedBoard int, pickedCell int) (bool,error) {

    if !g.Boards[pickedBoard].CanPlay {
        return false, fmt.Errorf("This board can't be played on this turn.")
    }
    if g.Boards[pickedBoard].XMap[pickedCell] {
        return false, fmt.Errorf("X already played in this cell.")
    }
    if g.Boards[pickedBoard].OMap[pickedCell] {
        return false, fmt.Errorf("O already played in this cell.")
    }

    if g.Turn % 2 == 0 {
        g.Boards[pickedBoard].XMap[pickedCell] = true
    } else {
        g.Boards[pickedBoard].OMap[pickedCell] = true
    }


    freePlace := g.Boards[pickedCell].isFull() || g.Boards[pickedCell].XMap.isTickTacToe()  || g.Boards[pickedCell].OMap.isTickTacToe()
    c := 0
    for i := 0; i < len(g.Boards); i++ {
        g.Boards[i].CanPlay = !(g.Boards[pickedCell].isFull() || g.Boards[i].XMap.isTickTacToe() || g.Boards[i].OMap.isTickTacToe()) && ((i==pickedCell) || freePlace)
        if g.Boards[i].CanPlay {
            c++
        }
    } 
    if c == 0 {
        return true, nil
    }

    g.XMap[pickedBoard] = g.Boards[pickedBoard].XMap.isTickTacToe()
    g.OMap[pickedBoard] = g.Boards[pickedBoard].OMap.isTickTacToe()


    isWinning := g.XMap.isTickTacToe() || g.OMap.isTickTacToe()



    g.Turn++
    if isWinning {
        fmt.Print("someone won\n")
        g.print()
    }
    return isWinning, nil

}

func (g *Game) ValidMoves() [][2]int {
    moves := make([][2]int, 0, 81) //max we can ever have is 81

    for i := 0; i < len(g.Boards); i++ {
        if !g.Boards[i].CanPlay {
                continue
        }
        for j := 0; j < len(g.Boards); j++ {
            if g.Boards[i].XMap[j] || g.Boards[i].OMap[j] {
                continue
            }
            moves = append(moves, [2]int{i,j})
        }
    }
    return moves
}


func main() {
    g := NewGame()
    

    for {
        aiBoard := rand.Intn(9)
        aiCell := rand.Intn(9)

        result, err := g.TakeTurn(aiBoard, aiCell)
        if err != nil {
            //fmt.Printf(err.Error())
        }else {
            fmt.Printf("[%v, %v]\n", aiBoard, aiCell)

        }
        if result {
            fmt.Println("Game is concluded")
            break
        }

    }
    g.print()

}
