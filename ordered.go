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

func orderedAndIntervaled(ints []int) bool {
	if len(ints) <= 1 {
		return true;
	}

	last := ints[0]
	interval := ints[1] - ints[0]
	for indy, i := range ints {
		if indy == 0 {
			continue
		}

        if i < last || i - last != interval {
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
	fmt.Println(orderedAndIntervaled([]int{1, 2, 5, 8})) // returns false
	fmt.Println(orderedAndIntervaled([]int{8, 6, 4, 2})) // returns false
	fmt.Println(orderedAndIntervaled([]int{2, 4, 6, 8})) // returns true
	fmt.Println(orderedAndIntervaled([]int{5, 10, 15, 20})) // returns true
}
