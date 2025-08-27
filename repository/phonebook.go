package repository

import (
	"asean-phonebook/model"
	"errors"
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
			pb.Contacts = append(pb.Contacts[:i], pb.Contacts[i+1:]...)
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
		// make room by appending nil then shifting
		pb.Contacts = append(pb.Contacts, nil)
		copy(pb.Contacts[index+1:], pb.Contacts[index:])
	case "delete":
		// shift left to remove
		copy(pb.Contacts[index:], pb.Contacts[index+1:])
		pb.Contacts = pb.Contacts[:len(pb.Contacts)-1]
	}
}

func (pb *Phonebook) PrintContactsFromCountryCodes(countryCodes []int) []*model.Person {
	var filtered []*model.Person
	lookup := make(map[int]struct{})
	for _, code := range countryCodes {
		lookup[code] = struct{}{}
	}
	for _, c := range pb.Contacts {
		if _, ok := lookup[c.GetCountryCode()]; ok {
			filtered = append(filtered, c)
		}
	}
	return filtered
}
