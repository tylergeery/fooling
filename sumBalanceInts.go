package main

import (
	"fmt"
)

func findBalanceIndex(ints []int) int {
    i, leftToRightSum, rightToLeftSum, length := 0, 0, 0, len(ints) - 1

    // edge cases where balance is not well-defined
    if length < 2 {
        return -1
    }

    for ; i < length; i++ {
        leftToRightSum += ints[i]
    }

    for ; i > 0; i-- {
        rightToLeftSum += ints[i]
        leftToRightSum -= ints[i - 1]

        if leftToRightSum == rightToLeftSum {
            return i - 1
        }
    }

    return -1
}

func main() {
    fmt.Println(findBalanceIndex([]int{})) // returns -1 for empty
    fmt.Println(findBalanceIndex([]int{1,40,1})) // returns 1, because we zero index like civilized people
    fmt.Println(findBalanceIndex([]int{3, 4, 2, -4, 100, 6, -1})) // returns 4. Left side is [3, 4, 2, -4] which equals 5, and right side is [6, -1] which also equals 5
    fmt.Println(findBalanceIndex([]int{3, 4, 8, 1, 5})) // returns -1
}
