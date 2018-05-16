package main

import (
	"fmt"
)

func double(ints []int) []int {
    for i, l := 0, len(ints); i < l; i++ {
        ints[i] *= 2
    }

    return ints
}

func applyFunc(ints []int, f func(int) int) []int {
    for i, l := 0, len(ints); i < l; i++ {
        ints[i] = f(ints[i])
    }

    return ints
}

func main() {
    fmt.Println(double([]int{1, 2, 3, 4})) // returns [2,4,6,8]
    fmt.Println(double([]int{5, 20})) // returns [10, 40]

    bar := func (i int) int {
        return i *3
    }
    fmt.Println(applyFunc([]int{1,2,3,4}, bar)) // returns [3, 6, 9, 12]

    bizz := func (i int) int {
        if (i % 2 == 0) {
            return 2
        } else {
            return 1
        }
    }
    fmt.Println(applyFunc([]int{1,2,3,4}, bizz)) // returns [1, 2, 1, 2]
}
