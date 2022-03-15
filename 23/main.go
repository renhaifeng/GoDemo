package main

import "fmt"

func modifyArray(arr [3][]string) [3][]string {
	arr[1] = []string{"d", "e", "p"}
	return arr
}

func modifyArray1(arr [3][]string) [3][]string {
	arr[0][0] = "d1"
	return arr
}

func main() {
	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}

	complexArray2 := modifyArray(complexArray1)

	complexArray3 := modifyArray1(complexArray1)

	fmt.Printf("val:%#v\n", complexArray1)
	fmt.Printf("val:%#v\n", complexArray2)
	fmt.Printf("val:%#v\n", complexArray3)
}
