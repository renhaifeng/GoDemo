package main

import "fmt"

func main() {
	ch1 := make(chan int, 1) // 缓存长度为 1

	ch1 <- 10

	x := <-ch1

	fmt.Println(x)

	fmt.Println("len:", len(ch1))

	fmt.Println("cap:", cap(ch1))

	close(ch1)
}
