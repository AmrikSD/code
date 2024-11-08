# UTTT (Ultimate Tic Tac Toe)

We can think of UTTT as a grid of 81 squares (3*3) * (3*3), a 3 by 3 of tic tac toe boards.

This means, we could encode each move similarly to [PGN Notation](https://en.wikipedia.org/wiki/Portable_Game_Notation) in chess.
In UTTT, since we're only picking the square we can just repeat a number from 0-80 (index of the square)

The point of this program is to "solve" UTTT, by outputting nested directories of the form

```sh
├── 0
├── 1
├── ...
└── 33
    ├── 12
    │   ├── 3
    │   ├── ...
    │   └── 4
    └── 13
        ├── 0
        └── 1
├── 80
└── 79
```

Eventually these will "terminate" to a file containing the output of the game, so you can `cat 33/12/3` and the output will be something like `X Wins`

## Usage

./uttt <path to output>
