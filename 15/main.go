package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("hello ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("main.")
}
