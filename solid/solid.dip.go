package main

import "fmt"

/*
Dependency inversion principle
HLM should not depend on LLM
Both should depend on abstractions
*/

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}
type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range rs.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}
	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations, Info{parent, Parent, child})
	rs.relations = append(rs.relations, Info{child, Child, parent})
}

type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate(){
	for _, p := range r.browser.FindAllChildrenOf("John"{
		fmt.Println("J",p.name)
	})
}