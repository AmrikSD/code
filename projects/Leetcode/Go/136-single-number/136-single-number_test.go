package main

import (
    "testing"
)

type Test struct {
    explaination string
    input []int
    expected int
}
var Tests = []Test{
    {
        explaination: "case 1 in LC",
        input: []int{2,2,1},
        expected: 1,
    },
    {
        explaination: "case 2 in LC",
        input: []int{4,1,2,1,2},
        expected: 4,
    },
    {
        explaination: "case 3 in LC",
        input: []int{1},
        expected: 1,
    },
}

func TestLargestAltitude(t *testing.T) {

    for _, tt := range Tests {
        t.Run(tt.explaination, func(t *testing.T){
            result := singleNumber(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, expected %v for test %s", result, tt.expected, tt.explaination)
            }
        })
    }
}

