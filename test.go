package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

const limit = 100000
const maxNumber = 10000000

func genRandom() []int {
	randInts := make([]int, limit)

	for i := 0; i < limit; i++ {
		randInts[i] = rand.Intn(int(maxNumber))
	}
	
	return randInts
}

func genRandomStrings(randInts []int) []string {
	randStrings := make([]string, len(randInts))
	
	for i := 0; i < len(randInts); i++ {
		randStrings[i] = strconv.Itoa(randInts[i])
	}
	
	return randStrings
}

func digitSumRecursive(x int) int {
	if x <= 1 {
		return x
	}

	return (x % 10) + digitSumRecursive(x/10)
}

func digitSumLoop(x int) int {
	var sum int
	for x >= 1 {
		sum += x % 10
		x = x / 10
	}
	return sum
}

func digitSumCharCodesLoop(stringNumber string) int {
	var sum int
	for _, digit := range stringNumber {
		sum += int(digit - '0')
	}
	
	return sum
}

func timer() {
	rand.Seed(time.Now().UTC().UnixNano())
	randInts := genRandom()
	randStrings := genRandomStrings(randInts)

	start := time.Now()
	for _, i := range randInts {
		digitSumRecursive(i)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Finished digitSumRecursive, Took %d nanoseconds\n", elapsed.Nanoseconds())

	start = time.Now()
	for _, j := range randInts {
		digitSumLoop(j)
	}
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("Finished digitSumLoop, Took %d nanoseconds\n", elapsed.Nanoseconds())

	start = time.Now()
	for _, k := range randStrings {
		digitSumCharCodesLoop(k)
	}
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("Finished digitSumCharCodesLoop, Took %d nanoseconds\n", elapsed.Nanoseconds())
}

func test() {
	randInts := genRandom()
	randStrings := genRandomStrings(randInts)

	for i := 0; i < len(randInts); i++ {
		a := digitSumRecursive(randInts[i])
		b := digitSumLoop(randInts[i])
		c := digitSumCharCodesLoop(randStrings[i])

		if (a != b || a != c) {
			fmt.Printf("Solutions don't match %d %d %d\n", a, b, c)
		}
	}
}

func main() {
	timer()
}

