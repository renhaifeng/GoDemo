package main

import "fmt"

func main() {
	// new 用来初始化值类型，返回的是指针
	// make 用来初始化引用类型(slice、map、channel），返回的是类型本身

	a := new(int)

	b := a

	*b = 2

	fmt.Println(a, b)

	m1 := make([]int, 3)

	m1[0] = 1

	m2 := m1

	fmt.Println(m1, m2)

	fmt.Printf("%p\n", m1)

	fmt.Printf("%p\n", m2)

	// 类型断言必须是一个方法类型
	v, ok := interface{}(m2).([]int)

	fmt.Println(v, ok)
}
