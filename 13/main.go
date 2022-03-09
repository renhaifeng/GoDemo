package main

import (
	"fmt"
	"os"
)

// 学员信息管理系统

type Class struct {
	Class map[int]*Student
}

type Student struct {
	id   int
	name string
}

// 构造函数
func newStudent(id int, name string) *Student {
	return &Student{
		id:   id,
		name: name,
	}
}

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员")
	fmt.Println("2. 编辑学员")
	fmt.Println("3. 学员列表")
	fmt.Println("4. 退出系统")
}

func exitSystem() {
	fmt.Println("已退出...")
	os.Exit(0)
}

func (c *Class) addStudent() {
	var id int
	var name string
	fmt.Println("请输入学员编号")
	_, err := fmt.Scanf("%d\n", &id)
	fmt.Println("请输入学员姓名")
	_, err = fmt.Scanf("%s\n", &name)
	if err != nil {
		fmt.Println("添加学员错误！")
		return
	}
	_, ok := c.Class[id]
	if ok {
		fmt.Println("学员信息已存在！")
		return
	}
	student := newStudent(id, name)
	c.Class[id] = student
	fmt.Println("保存成功")
}

func (c *Class) editStudent() {
	fmt.Println("编辑学员")
}

func (c *Class) listStudent() {
	fmt.Printf("%s\t%s\t\n", "ID", "姓名")

	for _, stu := range c.Class {
		fmt.Printf("%d\t%s\t\n", stu.id, stu.name)
	}
}

func main() {
	c := &Class{}
	c.Class = make(map[int]*Student, 20)

	for {
		// 打印系统菜单
		showMenu()

		// 获取用户输入
		var input int
		_, err := fmt.Scanf("%d\n", &input)
		if err != nil {
			fmt.Println("输入错误，请重新输入")
			continue
		}

		switch input {
		case 1:
			c.addStudent()
		case 2:
			c.editStudent()
		case 3:
			c.listStudent()
		case 4:
			exitSystem()
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}
