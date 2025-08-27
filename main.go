package main

import (
	"asean-phonebook/model"
	"asean-phonebook/repository"
	"fmt"
)

func main() {
	p1 := model.NewPerson(1, "Alice", "Smith", 1, 123, "1111111", "Engineer", "F")
	p2 := model.NewPerson(2, "Bob", "Adams", 44, 456, "2222222", "Teacher", "M")
	p3 := model.NewPerson(3, "Charlie", "Smith", 1, 123, "3333333", "Doctor", "M")

	pb := &repository.Phonebook{Contacts: []*model.Person{}}

	pb.Insert(p1)
	pb.Insert(p2)
	pb.Insert(p3)

	fmt.Println("Contacts after insertion (sorted by last, first name):")
	for _, c := range pb.Contacts {
		fmt.Printf("%d: %s, %s, Phone: %s\n", c.ID, c.LName, c.FName, c.GetPhoneNumber())
	}

	pb.DeleteContact(2)
	fmt.Println("\nContacts after deleting ID 2:")
	for _, c := range pb.Contacts {
		fmt.Printf("%d: %s, %s\n", c.ID, c.LName, c.FName)
	}

	filtered := pb.PrintContactsFromCountryCodes([]int{1})
	fmt.Println("\nContacts with country code 1:")
	for _, c := range filtered {
		fmt.Printf("%d: %s, %s\n", c.ID, c.LName, c.FName)
	}
}
