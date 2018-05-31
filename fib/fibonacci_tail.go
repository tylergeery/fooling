package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib_tail(n, last, curr int) int {
	if n == 0 {
		return curr
	}

	return fib_tail(n-1, curr, last+curr)
}

// func fib_recursive(n int) int {
// 	if n <= 2 {
// 		return 1
// 	}
//
// 	return fib_recursive(n-1) + fib_recursive(n-2)
// }

func fib_iter(n int) int {
	n1, curr := 1, 1
	if n <= 2 {
		return 1
	}

	for i := 2; i < n; i++ {
		n1, curr = curr, curr+n1
	}

	return curr
}

func main() {
	nth, _ := strconv.Atoi(os.Args[1])

	fmt.Println(fib_tail(nth, 1, 0))
	// fmt.Println(fib_recursive(nth))
	fmt.Println(fib_iter(nth))
}
