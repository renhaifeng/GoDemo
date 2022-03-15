package main

import (
	"fmt"
)

func main() {

	CommandLine := "aa"

	val, ok := interface{}(CommandLine).(int)

	fmt.Println(val, ok)
}
