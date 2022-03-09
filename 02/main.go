package main

import "fmt"

func sayHi(hello string) func(string) string {
	return func(word string) string {
		return hello + " " + word
	}
}

func main() {
	r := sayHi("hello")
	fmt.Println(r("world!!!"))
}
