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

func main() {
	intExpand()

	chineseExpand()

	stringExpand()
	
	stringSlice := []int{1}

	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", stringSlice, len(stringSlice), cap(stringSlice), stringSlice)

	stringSlice = append(stringSlice, 2, 3, 4)

	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", stringSlice, len(stringSlice), cap(stringSlice), stringSlice)

}
