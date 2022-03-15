package main

import (
	"flag"
	. "fmt"
)

var a1 int

func main() {
	var name string // [1]

	flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]

	flag.Parse()

	Printf("Hello, %v!\n", name)

	a1 = a()

	//a1, a2, a3 := b()

	Println(a1)
}

func a() int {
	return 1
}

func b() (int, int, int) {
	return 3, 4, 5
}
