package repository

import (
	"asean-phonebook/model"
	"errors"
	"strings"
)

type Phonebook struct {
	Contacts []*model.Person
}

func NewPhonebook() *Phonebook {
	return &Phonebook{Contacts: []*model.Person{}}
}

func (pb *Phonebook) GetContactAtIndex(index int) (*model.Person, error) {
	if index < 0 || index >= len(pb.Contacts) {
		return nil, errors.New("index out of range")
	}
	return pb.Contacts[index], nil
}

func (pb *Phonebook) GetContact(id int) (*model.Person, error) {
	for _, contact := range pb.Contacts {
		if contact.GetID() == id {
			return contact, nil
		}
	}
	return nil, errors.New("contact not found")
}

func (pb *Phonebook) IsEmpty() bool {
	return len(pb.Contacts) == 0
}

func (pb *Phonebook) DeleteContact(id int) error {
	for i, contact := range pb.Contacts {
		if contact.GetID() == id {
			pb.adjustPhonebook(i, "delete")
			return nil
		}
	}
	return errors.New("contact not found")
}

func (pb *Phonebook) Insert(p *model.Person) {
	if len(pb.Contacts) == 0 {
		pb.Contacts = append(pb.Contacts, p)
		return
	}
	index := pb.findIndexInsertion(p)
	pb.adjustPhonebook(index, "insert")
	pb.Contacts[index] = p

}

func (pb *Phonebook) findIndexInsertion(p *model.Person) int {
	high := len(pb.Contacts) - 1
	low := 0

	for low <= high {
		mid := (high + low) / 2
		if pb.Contacts[mid].CompareTo(p) < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func (pb *Phonebook) adjustPhonebook(index int, direction string) {
	switch direction {
	case "insert":
		// to right
		pb.Contacts = append(pb.Contacts, nil)
		copy(pb.Contacts[index+1:], pb.Contacts[index:])
	case "delete":
		// to left
		copy(pb.Contacts[index:], pb.Contacts[index+1:])
		pb.Contacts = pb.Contacts[:len(pb.Contacts)-1]
	}
}

func (pb *Phonebook) PrintContactsFromCountryCodes(selectedCountryCodes []int) []string {
	var filtered []string
	lookup := make(map[int]struct{})
	for _, code := range selectedCountryCodes {
		lookup[code] = struct{}{}
	}
	for _, c := range pb.Contacts {
		if _, ok := lookup[c.GetCountryCode()]; ok {
			filtered = append(filtered, c.GetPersonDetails())
		}
	}
	return filtered
}

func (pb *Phonebook) GetSurnames(selectedSurname string) []string {
	var surnames []string
	for _, c := range pb.Contacts {
		if strings.EqualFold(c.GetLastName(), selectedSurname) {
			surnames = append(surnames, c.GetPersonDetails())
		}
	}
	return surnames
}
