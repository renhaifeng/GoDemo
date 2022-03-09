package main

import "fmt"

func funcA() {
	fmt.Println("funcA")
}

func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("funcB")
		}
	}()
	panic("panic funcB")
}

func funcC() {
	fmt.Println("funcC")
}

func main() {
	funcA()
	funcB()
	funcC()
}
