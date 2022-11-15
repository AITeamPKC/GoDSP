package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAdress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func main() {
	john := Person{"John", &Address{"123", "London", "UK"}, []string{"C", "M", "S"}}

	jane := john.DeepCopy()

	jane.Name = "J"
	jane.Address.StreetAdress = "321"
	jane.Friends = append(jane.Friends, "J")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
