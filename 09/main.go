package main

import "fmt"

func main() {
	// 打印斐波那契数列
	var n1, n2 uint

	for i := 1; i <= 100; i++ {
		if n1 == 0 {
			n1 = 1
		} else {
			n1, n2 = n2, n1+n2
		}
		fmt.Printf("result%d = %d\n", i, n1+n2)
	}
}
