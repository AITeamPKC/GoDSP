package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

func main() {
	p2 := NewPerson("Jane", 21)
	p2.Age = 30
	fmt.Println(p2)
}
