package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(num int) int {
		return base + num
	}

	sub := func(num int) int {
		return base - num
	}

	return add, sub
}

func main() {
	addFunc, subFunc := calc(10)

	fmt.Println(addFunc(20), subFunc(5))
	fmt.Println(addFunc(10), subFunc(52))
	fmt.Println(addFunc(3), subFunc(2))
}
