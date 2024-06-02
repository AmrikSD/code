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
        input: []int{-5,1,5,0,-7},
        expected: 1,
    },
    {
        explaination: "case 2 in LC",
        input: []int{-4,-3,-2,-1,4,3,2},
        expected: 0,
    },
}

func TestLargestAltitude(t *testing.T) {

    for _, tt := range Tests {
        t.Run(tt.explaination, func(t *testing.T){
            result := largestAltitude(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, expected %v for test %s", result, tt.expected, tt.explaination)
            }
        })
    }
}

