package main

import (
	"fmt"
	"time"
)

func main() {

	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		close(intChan)
		fmt.Println(2)
	})
	fmt.Println(1)
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}
