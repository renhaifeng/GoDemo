package main

import "fmt"

func main() {

	var a = []int{1, 2, 3, 4, 5}

	fmt.Printf("%#v %p %d %d\n", a, a, len(a), cap(a))

	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}
