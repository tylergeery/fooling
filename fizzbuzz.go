package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		help()
		return
	}

	i := 1;
	length, err := strconv.Atoi(os.Args[1])

	if err != nil {
		help()
		return
	}

	for ; i <= length; i++ {
		if i % 3 == 0 {
			fmt.Printf("Fizz")
		}

		if i % 5 == 0 {
			fmt.Printf("Buzz")
		}

		if i % 5 != 0 && i % 3 != 0 {
			fmt.Printf(strconv.Itoa(i))
		}

		fmt.Println("")
	}
}

func help() {
	fmt.Printf("Usage: %s <integer>\n", os.Args[0])
}
