package main

import "fmt"

type operate func(x, y int) int

func calc(x, y int, op operate) (int, error) {
	return op(x, y), nil
}

func main() {
	var op operate

	op = func(x, y int) int {
		return x + y
	}

	fmt.Println(calc(1, 2, op))

	op = func(x, y int) int {
		return x - y
	}

	fmt.Println(calc(1, 2, op))

	op = func(x, y int) int {
		return x * y
	}

	fmt.Println(calc(1, 2, op))

	op = func(x, y int) int {
		return x / y
	}

	fmt.Println(calc(1, 2, op))
}
