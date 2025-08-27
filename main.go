package main

import (
	"asean-phonebook/model"
)

func main() {
	p1 := &model.Person{Name: "Alice"}
	p2 := &model.Person{Name: "Bob"}
	p3 := &model.Person{Name: "Charlie"}

	pb := &Phonebook{}
	pb.Insert(p2)
	pb.Insert(p1)
	pb.Insert(p3)
}
