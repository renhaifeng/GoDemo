package main

import "fmt"

func main() {
	v := (*[3]int)(nil)
	fmt.Printf("v:%#v len:%d cap:%d\n", v, len(v), cap(v))

	var k = 9
	for k = range []int{} {

	}
	fmt.Println(k)

	for k = 0; k < 3; k++ {

	}
	fmt.Println(k)

	// [3]int 的指针强转为 nil，不能进行取值
	for k = range (*[3]int)(nil) {

	}
	fmt.Println(k)
}
