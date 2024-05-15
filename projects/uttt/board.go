package main

import (
	"errors"
)

type Board struct {
    Xs [9] bool
    Os [9] bool
}

type Game struct {
    move int
    canPick [9] bool
    board Board 
    miniGames [9]MiniGame
}

type MiniGame struct {
    board Board
}

func NewGame() Game {
    var b = Game{}

    for i := 0; i < len(b.canPick); i++ {
        b.canPick[i] = true
    }

    return b
}

func (g *Game) MakeMove(board int, cell int) (*Game, error) {
    if g.isOccupied(board, cell) {
        return g, errors.New("You can't play in that square, it's occupied")
    }
    if !g.canPick[board] {
        return g, errors.New("You may not play in that quadrant on this turn")
    }

    if g.move % 2 == 0 {
        g.miniGames[board].board.Xs[cell] = true
        won, _ := g.miniGames[board].board.IsTicTacToe()
        if won {
            g.board.Xs[board] = true
        }
    } else {
        g.miniGames[board].board.Os[cell] = true
        won, _ := g.miniGames[board].board.IsTicTacToe()
        if won {
            g.board.Os[board] = true
        }
    }

    g.move++ 
    g.setNextCanPlay(cell)
    return g, nil
}

func (g *Game) isOccupied(board int, cell int) bool {
    return g.miniGames[board].board.Os[cell] || g.miniGames[board].board.Xs[cell]
}

func (g *Game) setNextCanPlay(board int){
    // if the section is already won, then let them play anywhere.
    if g.board.Xs[board] || g.board.Os[board] {
        for i := 0; i < len(g.canPick); i++ {
            g.canPick[i] = true
        }
        return
    }
    // Otherwise, set only the section they can play in
    for i := 0; i < len(g.canPick); i++ {
        g.canPick[i] = false
    }
    g.canPick[board] = true
}


func (b *Board) IsTicTacToe() (bool, rune) {
    if isTicTacToe(b.Os){
        return true, 'O'
    }
    if isTicTacToe(b.Xs){
        return true, 'X'
    }
    return false, 'Q' //TODO: don't return Q here this is dumb
}

func isTicTacToe(pieces [9]bool) bool {
    var wins = [8][9]bool{
        {
            true, true, true,
            false, false, false,
            false, false, false,
        },
        {
            false, false, false,
            true, true, true,
            false, false, false,
        },
        {
            false, false, false,
            false, false, false,
            true, true, true,
        },
        {
            true, false, false,
            true, false, false,
            true, false, false,
        },
        {
            false, true, false,
            false, true, false,
            false, true, false,
        },
        {
            false, false, true,
            false, false, true,
            false, false, true,
        },
        {
            true, false, false,
            false, true, false,
            false, false, true,
        },
        {
            false, false, true,
            false, true, false,
            true, false, false,
        },
    }

    for _, win :=  range wins {
        if win == pieces {
            return true
        }
    }
    return false;
}
