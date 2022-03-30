package main

import "fmt"

// 用100文钱买一百只鸡，其中公鸡，母鸡，小鸡都必须要有，问公鸡，母鸡，小鸡要买多少只刚好凑足100文钱

func main() {
	for g := 1; g < 20; g++ {
		for m := 1; m < 33; m++ {
			x := 100 - g - m
			if (x%3 == 0) && ((5*g + 3*m + x/3) == 100) {
				fmt.Printf("公鸡: %d只，母鸡：%d只，小鸡 %d只\n", g, m, x)
			}
		}
	}
}
