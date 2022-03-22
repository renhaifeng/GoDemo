package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
}

func main() {
	slice := make([]int, 5, 5)
	// 初始化一个长度为5，容量为5的 int 类型切片 [0, 0, 0, 0, 0]
	slice[0] = 1
	slice[1] = 2
	// 切片为引用类型 [1, 2, 0, 0, 0]
	// Go 提供的语法糖...，可以将 slice 传进可变函数，不会创建新的切片
	change(slice...)
	// 第一次 append 之后会扩容，创建一个新的切片，不会对原切片造成影响
	fmt.Println(slice)
	// 使用操作符 [i:j] 获得一个新切片 [1, 2]（长度为2，容量为5，和 slice 共用底层数组）
	change(slice[0:2]...)
	// 第二次 append 之后不会扩容，修改原切片 [1, 2, 3, 0, 0]
	fmt.Println(slice)
}
