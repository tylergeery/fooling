package fib

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func Fib_tail(n int, last, curr *big.Int) *big.Int {
	if n == 0 {
		return curr
	}

	return Fib_tail(n-1, curr, big.NewInt(0).Add(last, curr))
}

func Fib_iter(n int) *big.Int {
	tmp, n1, curr := big.NewInt(1), big.NewInt(1), big.NewInt(1)
	if n <= 2 {
		return big.NewInt(1)
	}

	for i := 2; i < n; i++ {
		tmp = curr
		curr = big.NewInt(0).Add(curr, n1)
		n1 = tmp
	}

	return curr
}

func main() {
	nth, _ := strconv.Atoi(os.Args[1])

	fmt.Println(Fib_tail(nth, big.NewInt(1), big.NewInt(0)))
	fmt.Println(Fib_iter(nth))
}
