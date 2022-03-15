package main

import "fmt"

type operate func(x, y int) int

func calc() (func(int, int) int, func(int, int) int, func(int, int) int, func(int, int) int) {
	add := func(x, y int) int {
		return x + y
	}

	sub := func(x, y int) int {
		return x - y
	}

	pow := func(x, y int) int {
		return x * y
	}

	div := func(x, y int) int {
		return x / y
	}

	return add, sub, pow, div
}

func main() {
	add, sub, pow, div := calc()

	fmt.Println(add(1, 2))

	fmt.Println(sub(1, 2))

	fmt.Println(pow(1, 2))

	fmt.Println(div(1, 2))
}
