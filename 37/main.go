package main

import (
	"fmt"
)

func main() {

	str1 := "abc"

	fmt.Println(len(str1))

	slice1 := make([]int, 3)

	fmt.Println(slice1, cap(slice1))

	//c := make(chan os.Signal)
	//signal.Notify(c)
	//fmt.Println("start..")
	//s := <-c
	//fmt.Println("End...", s)
}
