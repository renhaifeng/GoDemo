package main

import (
	"fmt"
	"strings"
)

func worldCount(str string) map[string]int {
	strSlice := strings.Split(str, " ")
	worldCount := make(map[string]int, 10)

	for _, word := range strSlice {
		worldCount[word]++
	}

	return worldCount
}

func worldCount1(str string) (worldCount map[string]int) {
	strSlice := strings.Split(str, " ")
	worldCount = make(map[string]int, 10)

	for _, word := range strSlice {
		worldCount[word]++
	}

	return
}

func main() {

	fmt.Println(worldCount("how do you do"))
	fmt.Println(worldCount1("what are you nong sha lei you"))
}
