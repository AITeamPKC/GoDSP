package main

type Bird struct {
	age int
}

func (b *Bird) Age() int       { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }
