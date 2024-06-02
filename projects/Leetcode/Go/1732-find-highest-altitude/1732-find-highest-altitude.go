package main

func largestAltitude(gain []int) int {
    max := 0
    curr := 0
    for _,v  := range gain {
        curr = curr + v
        if curr > max {
            max = curr
        }
    }
    return max
}
