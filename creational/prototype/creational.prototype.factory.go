package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var mainOffice = Employee{"", Address{0, "123", "London"}}
var auxOffice = Employee{"", Address{0, "66", "London"}}

func newEmployee(prt *Employee, name string, suite int) *Employee {
	result := prt.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewAuxOfficeEmployee("Jo", 100)
	jane := NewMainOfficeEmployee("Ja", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
