package main

import "fmt"

func intExpand() {
	//append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

func chineseExpand() {
	worldList := []string{
		"北京",
		"上海",
		"广州",
		"深圳",
		"河北",
		"河南",
		"山西",
		"天津",
		"黑龙江",
		"内蒙",
	}
	var chineseSlice []string
	for _, world := range worldList {
		chineseSlice = append(chineseSlice, world)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p \n", chineseSlice, len(chineseSlice), cap(chineseSlice), chineseSlice)
	}
}

func stringExpand() {
	stringList := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	var stringSlice []string
	for _, str := range stringList {
		stringSlice = append(stringSlice, str)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", stringSlice, len(stringSlice), cap(stringSlice), stringSlice)
	}
}

func byteExpand() {
	byteList := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var byteSlice []byte
	for _, b := range byteList {
		byteSlice = append(byteSlice, b)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", byteSlice, len(byteSlice), cap(byteSlice), byteSlice)
	}
}

func main() {
	intExpand()

	chineseExpand()

	stringExpand()

	byteExpand()

	stringSlice := []int{1}

	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", stringSlice, len(stringSlice), cap(stringSlice), stringSlice)

	stringSlice = append(stringSlice, 2, 3, 4)

	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", stringSlice, len(stringSlice), cap(stringSlice), stringSlice)

}
