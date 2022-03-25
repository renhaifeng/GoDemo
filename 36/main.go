package main

import "fmt"

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func main() {
	f := F(5)
	defer func() {
		fmt.Println(f()) // 8
	}()
	defer fmt.Println(f()) // 6 defer 后面的函数带参数，优先计算参数，结果存储在栈中，真正执行的时候取出
	i := f()               // 7
	fmt.Println(i)
}
