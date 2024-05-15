package main

import "fmt"

func main(){
    var b = NewGame();

    // Top Left
    b.MakeMove(0,0)
    b.MakeMove(0,8)
    b.MakeMove(8,0)
    b.MakeMove(0,7)
    b.MakeMove(7,0)
    b.MakeMove(0,6)

    // Top Middle
    b.MakeMove(6,1)
    b.MakeMove(1,8)
    b.MakeMove(8,1)
    b.MakeMove(1,7)
    b.MakeMove(7,1)
    b.MakeMove(1,6)

    b.MakeMove(6,2)
    b.MakeMove(2,8)
    b.MakeMove(8,2)
    b.MakeMove(2,7)
    b.MakeMove(7,2)
    b.MakeMove(2,6)



    fmt.Printf("%v\n",b)

    isTTT, whoWon := b.board.IsTicTacToe()
    fmt.Printf("%v, %c", isTTT, whoWon)
}
