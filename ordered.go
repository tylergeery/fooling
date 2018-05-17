package main

import (
	"fmt"
)

func ordered(ints []int) bool {
	last := -(int(^uint(0) >> 1)  - 1);

    for _, i := range ints {
        if i < last {
			return false
		}

		last = i
    }

    return true
}

func main() {
	fmt.Println(ordered([]int{8, 5, 2, 1})) // returns false
	fmt.Println(ordered([]int{1, 2, 5, 8})) // returns true

	// and for the extra credit
	fmt.Println(ordered([]int{8, 6, 4, 2})) // returns false
	fmt.Println(ordered([]int{2, 4, 6, 8})) // returns true
}
