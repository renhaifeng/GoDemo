package main

import "fmt"

//var wg sync.WaitGroup

func f1(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go f1(ch1)
	go f2(ch1, ch2)

	for item := range ch2 {
		fmt.Println(item)
	}
}
