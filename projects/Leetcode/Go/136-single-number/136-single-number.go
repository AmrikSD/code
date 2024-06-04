package main

func singleNumber(nums []int) int {

    c := nums[0]

    if len(nums) == 1 {
        return c
    }

    for i := 1; i < len(nums); i++ {
        c = c ^ nums[i]
    }

    return c
}
