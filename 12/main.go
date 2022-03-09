package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

type Animal struct {
	string
	int8
}

func main() {
	person := Person{
		"张三",
		18,
		"男",
	}

	fmt.Println(person.name, person.age, person.gender)

	animal := Animal{
		"熊猫",
		18,
	}

	fmt.Printf("animal: %#v \n", animal)
	fmt.Println(animal.string, animal.int8)
}
