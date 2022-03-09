package main

import (
	"fmt"
)

func main() {
	// 打印斐波那契数列
	for i := 1; i <= 100; i++ {
		result := fib(i)
		fmt.Printf("result%d = %d\n", i, result)
	}
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
