package main

import (
	"fmt"
	"runtime"
)

var info string

var m map[int]string = map[int]string{1: "A", 2: "B", 3: "C"}

func init() {
	fmt.Printf("Map: %v\n", m)

	info = fmt.Sprintf("OS: %s, Arch: %s", runtime.GOOS, runtime.GOARCH)
}

func main() {
	fmt.Println(info)
}
